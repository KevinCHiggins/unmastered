package controller

import (
	models "intive/unmastered/models"
	"net/http"
	"strconv"
	"strings"
)

type scores singleAndMultipleController

func (c scores) registerRoutes() {
	http.HandleFunc("/scores", c.handleScores)
	http.HandleFunc("/scores/", c.handleProject)
}

func (c scores) handleScores(w http.ResponseWriter, r *http.Request) {
	m, _ := models.GetScores()
	c.multipleTemplate.Execute(w, m)
}

func (c scores) handleProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path, "/scores/")[1])
	m, err := models.GetProject(id)
	if err != nil {

		errorTemplate.Execute(w, err)
		return
	}
	c.singleTemplate.Execute(w, m)
}
