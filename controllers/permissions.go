package controllers

import (
	models "intive/unmastered/models"
	"net/http"
	"strconv"
)

type createPermission controller
type PermissionContext struct {
	Contributors []models.Contributor
	Projects     []models.Project
}

func (c createPermission) registerRoutes() {
	http.HandleFunc("/create-permission", c.handleCreatePermission)
}

func (c createPermission) handleCreatePermission(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		v := getFormFieldValues(r, "contributor-radio-button", "project-radio-button", "type-radio-button")
		contributorId, err := strconv.Atoi(v[0])
		if err != nil {
			panic("Could not convert contributor ID string to integer")
		}
		projectId, err := strconv.Atoi(v[0])
		if err != nil {
			panic("Could not convert project ID string to integer")
		}
		if v[3] == "read" {
			models.CreateReadPermission(contributorId, projectId)
		} else if v[3] == "write" {
			models.CreateWritePermission(contributorId, projectId)
		} else {
			panic("Unexpected value for permission type")
		}
	} else {
		print("Dah")
		contributors, err := models.GetContributors()
		if err != nil {
			panic("")
		}
		for _, c := range contributors {
			print(c.Name)
		}
		projects, err := models.GetProjects()
		if err != nil {
			panic("")
		}
		for _, p := range projects {
			print(p.Moniker)
		}
		context := PermissionContext{
			Contributors: contributors,
			Projects:     projects,
		}

		c.template.Execute(w, context)
	}
}
