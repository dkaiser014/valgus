package main

import (
	"log"
	"net/http"
	"valgus/pkg/router"
)

// TODO's
// > First thing i want to do is running a http server using go `net/http` package. (Done)
// > Then I want to be able to render and .html file. (Done)
// > Once I'm able to render .html files I want to make a Router to handle all the request. (Done)
// > After that I would like to implement error handling for non-existing pages.
// > To finish off it will be nice to use Pug and Sass as transpilers.

func main() {
	router.Router()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
