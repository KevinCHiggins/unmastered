package model

import (
	"time"
)

type MediaUpload struct {
	Id           int
	ProjectId    int
	Type         int
	Title        string
	Notes        string
	UploadedDate time.Time
}
