package model

import (
	"errors"
	"time"
)

var contributors []Contributor

type Contributor struct {
	Id         int
	Name       string
	Email      string
	Password   string
	JoinedDate time.Time
}

func GetContributors() ([]Contributor, error) {
	return contributors, nil
}

func GetContributor(id int) (Contributor, error) {
	// should be a different database query, of course, but for now
	contribs, _ := GetContributors()
	for _, contributor := range contribs {
		if contributor.Id == id {
			return contributor, nil
		}
	}
	return Contributor{0, "", "", "", time.Now()}, errors.New(
		"ID does not match any contributors")
}

func GetContributorByCredentials(email, password string) (Contributor, error) {
	// should be a different database query, of course, but for now
	contribs, _ := GetContributors()
	for _, contributor := range contribs {
		if contributor.Email == email && contributor.Password == password {
			return contributor, nil
		}
	}
	return Contributor{0, "", "", "", time.Now()}, errors.New(
		"Credentials do not match any known contributor")
}

func CreateContributor(name, email, password string) {
	contributors = append(contributors,
		Contributor{len(contributors) + 1, name, email, password, time.Now()})
}
