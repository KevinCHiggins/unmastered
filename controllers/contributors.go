package controllers

import (
	models "intive/unmastered/models"
	"net/http"
)

type createContributor controller

func (c createContributor) registerRoutes() {
	http.HandleFunc("/create-contributor", c.handleCreateContributor)
}

func (c createContributor) handleCreateContributor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		v := getFormFieldValues(r, "name", "email", "password")
		models.CreateContributor(v[0], v[1], v[2])
	} else {
		m, _ := models.GetAnnouncements() // ?
		c.template.Execute(w, m)
	}
}
