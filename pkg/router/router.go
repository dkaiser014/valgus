package router

import (
	"net/http"
	"valgus/pkg/handlers"
)

// Router loads a specific page by its handler
func Router() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/about/", handlers.AboutHandler)
}
