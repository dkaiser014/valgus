package main

import (
	"fmt"
	"log"
	"net/http"
)

// TODO's
// > First thing i want to do is running a http server using go `net/http` package.
// > Then I want to be able to render and .html file.
// > Once I'm able to render .html files I want to make a Router to handle all the request.
// > After that I would like to implement error handling for non-existing pages.
// > To finish off it will be nice to use Pug and Sass as transpilers.

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
