package controller

import (
	"net/http"
)

type home controller

func (c home) registerRoutes() {
	http.HandleFunc("/home", c.handleHome)
}

func (c home) handleHome(w http.ResponseWriter, r *http.Request) {
	c.template.Execute(w, nil)
}
