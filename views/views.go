// package view contains functions for the various requests

package views

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	_ "fmt"
	"html/template"
	"io/ioutil"
	"models"
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
		"Quotos",
	}

	iTmpl := templates["index"]
	if err := iTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func ListQuotes(w http.ResponseWriter, r *http.Request) {
	// set response headers and return response
	w.Header().Set("Content-Type", "application/json")

	appengineContext := appengine.NewContext(r)
	q := models.LoadQuotes(appengineContext)

	lquotes, err := json.Marshal(q)

	if err != nil {
		panic(err)
	}

	w.Write(lquotes)
}

// request handler for posting quote to database
func SnippetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusOK)
		return
	}

	appengineContext := appengine.NewContext(r)
	var Q models.QuoteAux
	var status int

	// set response headers and return response
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &Q)

	if err != nil {
		panic(err)
	}

	result := Q.Validate()

	msg := models.MakeMessage()

	if len(result) == 0 {
		msg.SetType("success")
		msg.AddText(msg.GetType(), "Text successfully added!")

		// create new quote and save to database
		newQuote := models.NewQuote(Q)

		key := datastore.NewIncompleteKey(appengineContext, "Quote", nil)

		if _, err := datastore.Put(appengineContext, key, newQuote); err != nil {
			panic(err)
		}

		status = http.StatusOK

	} else {
		msg.SetType("error")
		msg.AddTextMap(result)
		status = http.StatusBadRequest
	}

	resp, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//appengineContext.Infof("%v", newQuote)

	w.WriteHeader(status)
	w.Write(resp)
}
