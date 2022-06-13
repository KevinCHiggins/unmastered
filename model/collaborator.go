package model

import (
	"time"
)

type Collaborator struct {
	Name       string
	Email      string
	Password   string
	JoinedDate time.Time
}
