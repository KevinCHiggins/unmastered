package main

import (
	controller "intive/unmastered/controllers"
	view "intive/unmastered/views"
	"net/http"
)

func main() {

	controller.ServePublicDir("/")
	controller.Setup(view.LoadTemplates())
	http.ListenAndServe(":8000", nil)
}
