package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type cat int

func (cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	host := strings.Split(req.Host, ":")

	fmt.Fprintln(w, "in /cat route")

	if host[1] == "8080" {
		io.WriteString(w, "using ServeMux")
	} else {
		io.WriteString(w, "using DefaultServerMux")
	}
}

type dog int

func (dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	host := strings.Split(req.Host, ":")

	fmt.Fprintln(w, "in /dog route")

	if host[1] == "8080" {
		io.WriteString(w, "using ServeMux")
	} else {
		io.WriteString(w, "using DefaultServerMux")
	}
}

// OPTION 1: use ServeMux of the http package
// https://golang.org/pkg/net/http/#ServeMux
func httpServeMux() {
	mux := http.NewServeMux()

	var c cat
	var d dog

	//Note:
	// https://golang.org/pkg/net/http/#ServeMux.Handle
	// func (mux *ServeMux) Handle(pattern string, handler Handler)
	// c & d implements the Hander interface
	// so they are instances of handlers
	mux.Handle("/cat", c)
	mux.Handle("/dog", d)

	http.ListenAndServe(":8080", mux)
}

// OPTION 2: use DefaultServerMux of the http package
func httpDefaultMux() {
	var c cat
	var d dog

	// https://golang.org/pkg/net/http/#Handle
	// func Handle(pattern string, handler Handler)
	http.Handle("/cat", c)
	http.Handle("/dog", d)

	//From Doc:
	// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux
	http.ListenAndServe(":8081", nil)
}

func main() {
	// comment and uncomment either function to see the difference
	// note that httpServeMux is listening on port 8080
	// while DefaultServerMux is listening on port 8081
	httpServeMux()

	// httpDefaultMux()
}
