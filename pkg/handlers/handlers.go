package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// Page represents the structure of the file
type Page struct {
	Title string
	Body  []byte
}

// LoadPage returns the title and the body(content)
// of a specific file
func LoadPage(title string) (*Page, error) {
	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// RenderTemplate parses the contents of a file
// and renders it to a template
func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles("web/views" + string(os.PathSeparator) + tmpl + ".html")
	t.Execute(w, p)
}

// IndexHandler handles the index route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	p, _ := LoadPage(title)
	RenderTemplate(w, "index", p)
}

// AboutHandler handles the about route
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path
	if r.URL.Path != "/about/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	p, _ := LoadPage(title)
	RenderTemplate(w, "about", p)
}

// ErrorHandler handles not found pages
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	title := "/error/"
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		p, _ := LoadPage(title)
		RenderTemplate(w, "error", p)
	}
}
