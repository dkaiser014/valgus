package handlers

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Page represents a structure for the content
// of the website
type Page struct {
	Title string
	Body  []byte
}

// LoadPage reads the file and returns the title and
// the body(content) of the file
func LoadPage(title string) (*Page, error) {
	filename := title
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Errored:", err)
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// RenderTemplate parses the file and writes the
// output to a template
func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// IndexHandler loads a page based on a URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := LoadPage(title)
	RenderTemplate(w, "index", p)
}

// AboutHandler loads a page based on a URL
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/about/"):]
	p, _ := LoadPage(title)
	RenderTemplate(w, "about", p)
}
