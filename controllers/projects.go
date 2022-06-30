package controller

import (
	models "intive/unmastered/models"
	"net/http"
	"strconv"
	"strings"
)

type projects singleAndMultipleController

func (c projects) registerRoutes() {
	http.HandleFunc("/projects", c.handleProjects)
	http.HandleFunc("/projects/", c.handleProject)
}

func (c projects) handleProjects(w http.ResponseWriter, r *http.Request) {
	m, _ := models.GetProjects()
	c.multipleTemplate.Execute(w, m)
}

func (c projects) handleProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path, "/projects/")[1])
	m, err := models.GetProject(id)
	if err != nil {

		errorTemplate.Execute(w, err)
		return
	}
	c.singleTemplate.Execute(w, m)
}
