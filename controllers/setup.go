package controller

import "html/template"

func Setup(templates map[string]*template.Template) {
	homeController = home{templates["home"]}
	homeController.registerRoutes()
	projectsController = projects{templates["project"], templates["projects"]}
	projectsController.registerRoutes()
	announcementsController = announcements{templates["announcements"]}
	announcementsController.registerRoutes()
	errorTemplate = templates["error"]

}
