package model

import (
	"fmt"
	"time"
)

type Project struct {
	Id             int
	Moniker        string
	RegisteredDate time.Time
}

var hardcodedProjects []Project

func GetProjects() ([]Project, error) {

	return hardcodedProjects, nil
}

func GetProject(id int) (Project, error) {
	projects, err := GetProjects()
	if err == nil {
		for _, p := range projects {
			if p.Id == id {
				return p, nil
			}
		}
		return Project{0, "", time.Now()}, fmt.Errorf("Project with Id %v not found", id)
	} else {
		return Project{0, "", time.Now()}, fmt.Errorf("Couldn't retrieve all projects")
	}
}
