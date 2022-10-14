package model

var readPermissions []ReadPermission

type ReadPermission struct {
	ProjectId     int
	ContributorId int
}

func CreateReadPermission(projectId, collaboratorId int) {
	readPermissions = append(readPermissions,
		ReadPermission{projectId, collaboratorId})
}
