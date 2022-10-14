package controllers

import (
	"html/template"
)

func Setup(templates map[string]*template.Template) {
	homeController = home{templates["home"]}
	homeController.registerRoutes()
	loginController = login{templates["login"]}
	loginController.registerRoutes()
	projectsController = projects{templates["project"], templates["projects"]}
	projectsController.registerRoutes()
	scoresController = scores{templates["score"], templates["scores"]}
	scoresController.registerRoutes()
	multitracksController = multitracks{templates["multitrack"], templates["multitracks"]}
	multitracksController.registerRoutes()
	samplePacksController = samplePacks{templates["sample-pack"], templates["sample-packs"]}
	samplePacksController.registerRoutes()
	announcementsController = announcements{templates["announcements"]}
	announcementsController.registerRoutes()
	createAnnouncementController = createAnnouncement{templates["create-announcement"]}
	createAnnouncementController.registerRoutes()
	createContributorController = createContributor{templates["create-contributor"]}
	createContributorController.registerRoutes()
	print(templates["create-permission"])
	createPermissionController = createPermission{templates["create-permission"]}
	createPermissionController.registerRoutes()
	errorTemplate = templates["error"]

	sessions = make(map[string]sessionState)

}
