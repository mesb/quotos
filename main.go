// main entry to the application
package main

import (
	"net/http"
	"views"
)


func init() {
	http.HandleFunc("/", views.IndexHandler)
	http.HandleFunc("/snippet", views.SnippetHandler)
	http.HandleFunc("/list", views.ListQuotes)

}
