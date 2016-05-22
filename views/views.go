// package view contains functions for the various requests

package views

import (
	"html/template"
	"net/http"
	"utils"
)

var templates map[string]template.Template

func init() {
	templates = make(map[string]template.Template)
	templates["index"] = *template.Must(template.ParseFiles(utils.BASE_TEMPLATE, utils.GetFilePath("index")))
}


// Index View
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title string
	}{
		"Kwata TV: Welcome",
	}

	iTmpl := templates["index"]
	if err := iTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
