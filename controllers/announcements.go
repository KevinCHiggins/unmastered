package controller

import (
	models "intive/unmastered/models"
	"net/http"
)

type announcements controller

func (c announcements) registerRoutes() {
	http.HandleFunc("/announcements", c.handleAnnouncements)
}

func (c announcements) handleAnnouncements(w http.ResponseWriter, r *http.Request) {
	m, _ := models.GetAnnouncements()
	c.template.Execute(w, m)
}
