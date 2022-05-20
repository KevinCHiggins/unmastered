package controller

import (
	"html/template"
	"intive/unmastered/viewmodel"
	"net/http"
	"regexp"
)

type mixes struct {
	template     *template.Template
	listTemplate *template.Template
}

func (m mixes) registerRoutes() {
	http.HandleFunc("/mixes", m.handleMixes)
	http.HandleFunc("/mixes/", m.handleMixes)
}

func (m mixes) handleMixes(w http.ResponseWriter, r *http.Request) {
	locationPattern, _ := regexp.Compile(`/mixes/(\w+)`)
	matches := locationPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		m.handleMix(w, r, matches[1])
	} else {
		vm := viewmodel.NewMixes()
		m.listTemplate.Execute(w, vm)
	}
}

func (m mixes) handleMix(w http.ResponseWriter, r *http.Request, location string) {
	mixesVm := viewmodel.NewMixes()
	for _, singleMixVm := range mixesVm.Mixes {
		if singleMixVm.MixLocation == location {
			m.template.Execute(w, singleMixVm)
		}
	}

}
