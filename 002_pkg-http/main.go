package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

// init as int just to show
// that anything implements the handler interface
// can serve as a handler
type cat int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	// cat implements the handler interface
	// so any variable defined as type cat
	// can be a handler
	var handler cat

	http.ListenAndServe(":8080", handler)
}

func (c cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// call function ParseForm first
	// so that form datas can be retrived
	// in req.Form
	err := req.ParseForm()

	if err != nil {
		log.Fatalln(err)
	}

	// set reqData struct
	reqData := struct {
		Method        string
		URL           *url.URL
		Form          map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	// set custom headers
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Header().Set("My-Custom-Header", "Hi there! Name's Levblanc!")

	// print directly to page
	fmt.Fprintln(w, "<h1 style='width: 1000px; margin: 0 auto;'>Request Info</h1>")
	// render with html template
	tpl.Execute(w, reqData)
}
