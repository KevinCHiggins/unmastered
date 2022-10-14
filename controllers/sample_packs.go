package controllers

import (
	models "intive/unmastered/models"
	"net/http"
	"strconv"
	"strings"
)

type samplePacks singleAndMultipleController

func (c samplePacks) registerRoutes() {
	http.HandleFunc("/sample-packs", c.handleSamplePacks)
	http.HandleFunc("/sample-packs/", c.handleSamplePack)
}

func (c samplePacks) handleSamplePacks(rw http.ResponseWriter, r *http.Request) {
	m, _ := models.GetSamplePacks()
	vm := buildViewModel(rw, r, m)
	c.multipleTemplate.Execute(rw, vm)
}

func (c samplePacks) handleSamplePack(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path, "/sample-packs/")[1])
	m, err := models.GetProject(id)
	if err != nil {

		errorTemplate.Execute(rw, err)
		return
	}
	vm := buildViewModel(rw, r, m)
	c.singleTemplate.Execute(rw, vm)
}
