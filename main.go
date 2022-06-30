package main

import (
	controllers "intive/unmastered/controllers"
	models "intive/unmastered/models"
	views "intive/unmastered/views"
	"net/http"
)

func main() {

	models.LoadTestData()
	controllers.ServePublicDir("/")
	controllers.Setup(views.LoadTemplates())
	http.ListenAndServe(":8000", nil)
}
