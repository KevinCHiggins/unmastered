package controller

import "html/template"

func Setup(templates map[string]*template.Template) {
	homeController = home{templates["home"]}
	homeController.registerRoutes()
	aboutController.registerRoutes()
	mixesController = mixes{templates["mix"], templates["mixes"]}
	mixesController.registerRoutes()
	remixesController = remixes{templates["remix"], templates["remixes"]}
	remixesController.registerRoutes()
}
