package controller

import (
	"fmt"
	"html/template"
)

func Setup(templates map[string]*template.Template) {
	homeController = home{templates["home"]}
	fmt.Println(home{templates["home"]})
	homeController.registerRoutes()
	projectsController = projects{templates["project"], templates["projects"]}
	projectsController.registerRoutes()
	scoresController = scores{templates["score"], templates["scores"]}
	scoresController.registerRoutes()
	announcementsController = announcements{templates["announcements"]}
	announcementsController.registerRoutes()
	errorTemplate = templates["error"]

}
