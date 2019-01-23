package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utr8")
	io.WriteString(w, "<img src='/public/geekCat.png' />")
}

func fileServer() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle(
		"/public/",
		http.StripPrefix("/public", http.FileServer(http.Dir("./assets"))),
	)

	http.ListenAndServe(":8080", nil)
}

func serveStaticSite() {
	// https://golang.org/pkg/net/http/#FileServer
	// From Doc:
	// As a special case, the returned file server redirects any request ending in "/index.html" to the same path, without the final "index.html".

	// This means:
	// any request to "/index.html" will be redirect to "/"
	// which also means:
	// if you request "/" and the directory served contains an "index.html"
	// the server will automatically serve this file
	log.Fatal(http.ListenAndServe(":8081", http.FileServer(http.Dir("."))))
}

func main() {
	fileServer()

	// serveStaticSite()
}
