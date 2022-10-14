package view

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
)

const (
	templateLoc = "views/"
	layoutLoc   = templateLoc + "layout/"
	contentLoc  = templateLoc + "pages/"
	specialLoc  = templateLoc + "special/" // pages w/o header or footer
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

	layout := template.Must(template.ParseFiles(layoutLoc + "layout.gohtml"))
	specialTemplates := inflateUsingDirContents(layout, specialLoc)

	template.Must(layout.ParseFiles(layoutLoc+"footer.gohtml",
		layoutLoc+"header.gohtml"))
	templates := inflateUsingDirContents(layout, contentLoc)

	for t := range specialTemplates {
		templates[t] = specialTemplates[t]
	}

	return templates
}

func inflateUsingDirContents(t *template.Template, location string) map[string]*template.Template {
	result := make(map[string]*template.Template)
	dirEntries := getDirEntries(location)
	for _, dirEntry := range dirEntries {
		templateFile, err := os.Open(location + dirEntry.Name())
		if err != nil {
			panic("File open error: " + err.Error())
		}
		content, err := io.ReadAll(templateFile)
		if err != nil {
			panic("File read error: " + err.Error())
		}
		templateFile.Close()
		pageTemplate := template.Must(t.Clone())
		_, err = pageTemplate.Parse(string(content))
		if err != nil {
			print(pageTemplate.Name)
			print(pageTemplate.Tree)
			fmt.Printf("Error detected at: %v", dirEntry.Name())
			panic("Parse error: " + err.Error())
		}
		dotIndex := strings.LastIndex(dirEntry.Name(), ".")
		templateName := dirEntry.Name()[:dotIndex]
		fmt.Println("Inflating template ", templateName)
		if templateName == "create-permission" {
			print(pageTemplate.Tree.Root.String())
		} else if templateName == "create-contributor" {
			print(pageTemplate.Tree.Root.String())
		}
		result[templateName] = pageTemplate
	}
	return result
}
