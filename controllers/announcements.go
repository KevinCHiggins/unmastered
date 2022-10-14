package controllers

import (
	model "intive/unmastered/models"
	models "intive/unmastered/models"
	"net/http"
)

type announcements controller
type createAnnouncement controller

func (c announcements) registerRoutes() {
	http.HandleFunc("/announcements", c.handleAnnouncements)
}

func (c announcements) handleAnnouncements(rw http.ResponseWriter, r *http.Request) {
	m, _ := models.GetAnnouncements()
	vm := buildViewModel(rw, r, m)
	c.template.Execute(rw, vm)
}

func (c createAnnouncement) registerRoutes() {
	http.HandleFunc("/create-announcement", c.handleCreateAnnouncement)
}

func (c createAnnouncement) handleCreateAnnouncement(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		values := getFormFieldValues(r, "title", "text")
		model.CreateAnnouncement(values[0], values[1])
	} else {
		m, _ := models.GetAnnouncements()
		vm := buildViewModel(rw, r, m)
		c.template.Execute(rw, vm)
	}
}
