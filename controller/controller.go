package controller

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

type controller struct {
	template *template.Template
}

var (
	homeController    home
	aboutController   about
	mixesController   mixes
	remixesController remixes
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
