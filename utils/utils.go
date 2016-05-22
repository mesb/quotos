// helper functions for the quotos application
// template loaders and other small utilities

package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// global constants
var (
	TEMPLATES_DIR string = filepath.Join(getCurrentDir(), "templates")
	BASE_TEMPLATE string = filepath.Join(TEMPLATES_DIR, "base.tmpl")
)

// return current directory
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

// Get file path
func GetFilePath(f string) string {
	return fmt.Sprintf("%s", TEMPLATES_DIR+"/"+f+".html")
}


// Page
type Page struct {
	Title string
	Templ *template.Template
}

func NewTemplate(t, f string) *Page {
	p := new(Page)
	p.Title = t
	p.Templ = template.Must(template.ParseFiles(BASE_TEMPLATE, GetFilePath(f)))

	return p
}
