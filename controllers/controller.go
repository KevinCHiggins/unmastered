package controllers

import (
	"fmt"
	"html/template"
	models "intive/unmastered/models"
	"io"
	"net/http"
	"os"
	"strings"
)

type viewModel struct {
	PageData interface{}
	User     *models.Contributor
}

type controller struct {
	template *template.Template
}

type singleAndMultipleController struct {
	singleTemplate   *template.Template
	multipleTemplate *template.Template
}

var (
	homeController               home
	loginController              login
	projectsController           projects
	announcementsController      announcements
	createAnnouncementController createAnnouncement
	createContributorController  createContributor
	createPermissionController   createPermission
	scoresController             scores
	multitracksController        multitracks
	samplePacksController        samplePacks
	errorTemplate                *template.Template
)

func ServePublicDir(path string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public/" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error: " + err.Error()))
		}
		println("Public" + " " + r.URL.Path)
		defer f.Close()
		var content []byte
		f.Read(content)
		println(content)
		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, "css"):
			{
				contentType = "text/css"

			}
		case strings.HasSuffix(r.URL.Path, "png"):
			{
				contentType = "image/png"
			}
		case strings.HasSuffix(r.URL.Path, "html"):
			{
				contentType = "text/html"
			}
		default:
			{
				contentType = "text/plain"
			}
		}
		println(contentType + " " + r.URL.Path + " ")
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)
	})
}

func getFormFieldValues(r *http.Request, keys ...string) []string {
	err := r.ParseForm()
	if err != nil {
		panic("")
	}
	if (len(r.Form)) != len(keys) {
		msg := fmt.Sprintf("Unexpected amount of fields: %v", len(r.Form))
		panic(msg)
	}
	var values []string
	for _, k := range keys {
		if !r.Form.Has(k) {
			msg := fmt.Sprintf("Unexpected field found: %v", k)
			panic(msg)
		}
		values = append(values, r.Form[k][0])
	}
	return values

}
