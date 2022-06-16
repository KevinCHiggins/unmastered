package model

import (
	"time"
)

type Project struct {
	Id             int
	Moniker        string
	RegisteredDate time.Time
}

func GetProjects() []Project {
	t, _ := time.Parse("02/01/06", "29/06/22")
	result := []Project{Project{1, "Kev and Soro", t}}
	return result
}
