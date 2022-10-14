package controllers

import (
	models "intive/unmastered/models"
	"net/http"
)

type login controller

func (c login) registerRoutes() {
	http.HandleFunc("/login", c.handleLogin)
}

func (c login) handleLogin(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.template.Execute(rw, nil)
	} else if r.Method == http.MethodPost {
		// should handle wrong field names
		v := getFormFieldValues(r, "email", "password")
		contributor, err := models.GetContributorByCredentials(v[0], v[1])
		if err != nil {
			print("Authentication failed")
			// TODO error page
			http.Redirect(rw, r, "/projects", http.StatusPermanentRedirect)
		} else {
			print("Authentication succeeded")
			sessionId := ensureSession(rw, r)
			session := sessions[sessionId]
			session.loggedIn = true // I hope this affects the actual object
			session.userId = contributor.Id
			sessions[sessionId] = session
			http.Redirect(rw, r, "/projects", http.StatusPermanentRedirect)
		}
	} else if r.Method == http.MethodDelete {
		requestSessionId, err := checkSession(r)
		if err == nil {
			deleteSession(requestSessionId)
		}
		http.Redirect(rw, r, "/projects", http.StatusPermanentRedirect)
	}
}
