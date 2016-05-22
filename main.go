// main entry to the application
package main

import (
	"net/http"
	"views"
)


func init() {
	http.HandleFunc("/", views.IndexHandler)
}
