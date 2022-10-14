package controllers

import (
	models "intive/unmastered/models"
	"net/http"
	"strconv"
	"strings"
)

type multitracks singleAndMultipleController

func (c multitracks) registerRoutes() {
	http.HandleFunc("/multitracks", c.handleMultitracks)
	http.HandleFunc("/multitracks/", c.handleMultitrack)
}

func (c multitracks) handleMultitracks(rw http.ResponseWriter, r *http.Request) {
	m, _ := models.GetMultitracks()
	vm := buildViewModel(rw, r, m)
	c.multipleTemplate.Execute(rw, vm)
}

func (c multitracks) handleMultitrack(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path, "/multitracks/")[1])
	m, err := models.GetProject(id)
	if err != nil {

		errorTemplate.Execute(rw, err)
		return
	}
	vm := buildViewModel(rw, r, m)
	c.multipleTemplate.Execute(rw, vm)
	c.singleTemplate.Execute(rw, vm)
}
