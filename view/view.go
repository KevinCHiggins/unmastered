package view

import (
	"html/template"
	"io"
	"os"
	"strings"
)

const (
	templateLoc = "templates/"
	layoutLoc   = templateLoc + "layout/"
	contentLoc  = templateLoc + "content/"
)

func getDirEntries(location string) []os.DirEntry {
	dir, err := os.Open(location)
	if err != nil {
		panic("Directory error: " + err.Error())
	}
	dirEntries, err := dir.ReadDir(-1)
	if err != nil {
		panic("Directory listing error: " + err.Error())
	}
	return dirEntries
}

func LoadTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(layoutLoc+"layout.gohtml",
		layoutLoc+"footer.gohtml", layoutLoc+"header.gohtml"))
	contentDirEntries := getDirEntries(contentLoc)
	for _, contentDirEntry := range contentDirEntries {
		contentFile, err := os.Open(contentLoc + contentDirEntry.Name())
		if err != nil {
			panic("File open error: " + err.Error())
		}
		content, err := io.ReadAll(contentFile)
		if err != nil {
			panic("File read error: " + err.Error())
		}
		contentFile.Close()
		pageTemplate := template.Must(layout.Clone())
		_, err = pageTemplate.Parse(string(content))
		if err != nil {
			panic("Parse error: " + err.Error())
		}
		dotIndex := strings.LastIndex(contentDirEntry.Name(), ".")
		templateName := contentDirEntry.Name()[:dotIndex]
		result[templateName] = pageTemplate
	}
	return result
}
