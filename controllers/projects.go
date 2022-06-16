package controller

import (
	"intive/unmastered/model"
	"net/http"
)

type projects singleAndMultipleController

func (c projects) registerRoutes() {
	http.HandleFunc("/projects", c.handleProjects)
}

func (c projects) handleProjects(w http.ResponseWriter, r *http.Request) {
	m := model.GetProjects()
	c.multipleTemplate.Execute(w, m)
}
