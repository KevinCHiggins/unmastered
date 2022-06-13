package model

import (
	"time"
)

type Project struct {
	Id             int
	Moniker        string
	RegisteredDate time.Time
}
