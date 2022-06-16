package model

import (
	"time"
)

type Collaborator struct {
	Id         int
	Name       string
	Email      string
	Password   string
	JoinedDate time.Time
}
