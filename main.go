package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// TODO's
// > First thing i want to do is running a http server using go `net/http` package.
// > Then I want to be able to render and .html file.
// > Once I'm able to render .html files I want to make a Router to handle all the request.
// > After that I would like to implement error handling for non-existing pages.
// > To finish off it will be nice to use Pug and Sass as transpilers.

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "index", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
