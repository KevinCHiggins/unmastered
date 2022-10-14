package controllers

import (
	"fmt"
	models "intive/unmastered/models"
	"net/http"
)

type sessionState struct {
	loggedIn bool
	userId   int
}

var currentSession int = 0

var sessions map[string]sessionState

func ensureSession(rw http.ResponseWriter, r *http.Request) string {
	requestSessionId, err := checkSession(r)

	// very basic, assuming that only error possible is missing cookie
	if err == nil {
		print("SessionId found in request")
		return requestSessionId
	}
	print("New session created")
	return resetSession(rw)
}

func getState(sessionId string) sessionState {
	return sessions[sessionId]
}

func checkSession(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil

}

func resetSession(rw http.ResponseWriter) string {
	newSessionId := newSession()
	newCookie := &http.Cookie{Name: "token", Value: newSessionId, Path: "/", Secure: true}
	http.SetCookie(rw, newCookie)
	return newSessionId
}

func deleteSession(sessionId string) {
	if _, exists := sessions[sessionId]; exists {
		delete(sessions, sessionId)
	}
}

func newSession() string {
	currentSession = currentSession + 1
	sessionId := fmt.Sprint(currentSession)
	sessions[sessionId] = sessionState{loggedIn: false, userId: 0}
	return sessionId
}

func getAuthenticatedContributorIfAny(rw http.ResponseWriter, r *http.Request) *models.Contributor {
	sessionId := ensureSession(rw, r)
	sessionState := getState(sessionId)
	userId := sessionState.userId
	user, err := models.GetContributor(userId)
	if err != nil {
		return nil
	}
	return &user
}

func buildViewModel(rw http.ResponseWriter, r *http.Request, PageData interface{}) viewModel {
	user := getAuthenticatedContributorIfAny(rw, r)
	return viewModel{PageData: PageData, User: user}
}
