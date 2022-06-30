package model

import (
	"fmt"
	"time"
)

type MediaUpload struct {
	Id             int
	ProjectId      int
	CollaboratorId int
	TypeId         int
	Title          string
	Notes          string
	Public         bool
	UploadedDate   time.Time
}

var hardcodedMediaUploads []MediaUpload

func GetMediaUploads() ([]MediaUpload, error) {

	return hardcodedMediaUploads, nil
}

func GetMediaUpload(id int) (MediaUpload, error) {
	mediaUploads, err := GetMediaUploads()
	if err == nil {
		for _, mu := range mediaUploads {
			if mu.Id == id {
				return mu, nil
			}
		}
		return MediaUpload{0, 0, 0, 0, "", "", false, time.Now()}, fmt.Errorf("Project with Id %v not found", id)
	} else {
		return MediaUpload{0, 0, 0, 0, "", "", false, time.Now()}, fmt.Errorf("Couldn't retrieve all projects")
	}
}
