package controller

import (
	"html/template"
	"intive/unmastered/viewmodel"
	"net/http"
	"regexp"
)

type remixes struct {
	template     *template.Template
	listTemplate *template.Template
}

func (rm remixes) registerRoutes() {
	http.HandleFunc("/remixes", rm.handleRemixes)
	http.HandleFunc("/remixes/", rm.handleRemixes)
}

func (rm remixes) handleRemixes(w http.ResponseWriter, r *http.Request) {
	// this doesn't work for locations with dashes
	locationPattern, _ := regexp.Compile(`/remixes/(\w+)`)
	matches := locationPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		rm.handleRemix(w, r, matches[1])
	} else {
		vm := viewmodel.NewRemixes()
		rm.listTemplate.Execute(w, vm)
	}
}

func (rm remixes) handleRemix(w http.ResponseWriter, r *http.Request, location string) {
	remixesVm := viewmodel.NewRemixes()
	for _, singleRemixVm := range remixesVm.Remixes {

		if singleRemixVm.RemixLocation == location {
			println("Found " + singleRemixVm.RemixLocation)
			rm.template.Execute(w, singleRemixVm)
		}
	}

}
