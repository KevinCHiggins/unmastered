package controller

import (
	"intive/unmastered/viewmodel"
	"net/http"
)

type about controller

func (a about) registerRoutes() {
	http.HandleFunc("/about", a.handleAbout)
}

func (a about) handleAbout(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewAbout()
	a.template.Execute(w, vm)
}
