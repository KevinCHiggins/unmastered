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

var hardcodedAnnouncements []Announcement

func GetAnnouncements() ([]Announcement, error) {
	return hardcodedAnnouncements, nil
}
