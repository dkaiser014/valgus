package router

import (
	"net/http"
	"valgus/pkg/handlers"
)

// Router loads a specific page by its handler
func Router() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/about/", handlers.AboutHandler)

	// Static asssets
	fs := http.FileServer(http.Dir("web/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/about/static/", http.StripPrefix("/about/static/", fs))
}
