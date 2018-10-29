// main entry to the application
package main

import (
	"net/http"
	"views"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func init() {
	router := mux.NewRouter()
	router.HandleFunc("/", views.IndexHandler)
	router.HandleFunc("/snippet", views.SnippetHandler)
	router.HandleFunc("/list", views.ListQuotes)

	handler := cors.Default().Handler(router)

	http.Handle("/", handler)
	handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(router)

}
