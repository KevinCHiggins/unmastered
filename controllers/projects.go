package controllers

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

func (c projects) handleProjects(rw http.ResponseWriter, r *http.Request) {
	m, _ := models.GetProjects()
	vm := buildViewModel(rw, r, m)
	c.multipleTemplate.Execute(rw, vm)
}

func (c projects) handleProject(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path, "/projects/")[1])
	m, err := models.GetProject(id)
	if err != nil {

		errorTemplate.Execute(rw, err)
		return
	}
	vm := buildViewModel(rw, r, m)
	c.singleTemplate.Execute(rw, vm)
}
