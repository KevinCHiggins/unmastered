package model

import (
	"time"
)

type MediaUpload struct {
	Id             int
	ProjectId      int
	CollaboratorId int
	TypeId         int
	Title          string
	Notes          string
	UploadedDate   time.Time
}
