package model

import (
	"time"
)

type Announcement struct {
	Id            int
	Title         string
	Text          string
	PublishedDate time.Time
}
