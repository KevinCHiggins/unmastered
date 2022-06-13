package controller

import (
	"intive/unmastered/viewmodel"
	"net/http"
)

type home controller

func (a home) registerRoutes() {
	http.HandleFunc("/home", a.handleHome)
}

func (a home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	a.template.Execute(w, vm)
}
