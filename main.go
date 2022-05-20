package main

import (
	"intive/unmastered/controller"
	"intive/unmastered/view"
	"net/http"
)

func main() {

	controller.ServePublicDir("/")
	controller.Setup(view.LoadTemplates())
	http.ListenAndServe(":8000", nil)
}
