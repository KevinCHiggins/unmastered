package controllers

import (
	models "intive/unmastered/models"
	"net/http"
	"strconv"
	"strings"
)

type scores singleAndMultipleController

func (c scores) registerRoutes() {
	http.HandleFunc("/scores", c.handleScores)
	http.HandleFunc("/scores/", c.handleScore)
}

func (c scores) handleScores(rw http.ResponseWriter, r *http.Request) {
	m, _ := models.GetScores()
	vm := buildViewModel(rw, r, m)
	c.multipleTemplate.Execute(rw, vm)
}

func (c scores) handleScore(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path, "/scores/")[1])
	m, err := models.GetProject(id)
	if err != nil {

		errorTemplate.Execute(rw, err)
		return
	}
	vm := buildViewModel(rw, r, m)
	c.singleTemplate.Execute(rw, vm)
}
