package model

var writePermissions []WritePermission

type WritePermission struct {
	ProjectId     int
	ContributorId int
}

func CreateWritePermission(projectId, contributorId int) {
	writePermissions = append(writePermissions,
		WritePermission{projectId, contributorId})
}
