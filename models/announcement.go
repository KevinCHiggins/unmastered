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

var announcements []Announcement

func GetAnnouncements() ([]Announcement, error) {
	return announcements, nil
}

func CreateAnnouncement(title, text string) {
	announcements = append(announcements,
		Announcement{len(announcements) + 1, title, text, time.Now()})
}
