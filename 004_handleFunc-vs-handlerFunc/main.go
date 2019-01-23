package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func c(w http.ResponseWriter, req *http.Request) {
	host := strings.Split(req.Host, ":")

	fmt.Fprintln(w, "in /cat route")

	if host[1] == "8080" {
		io.WriteString(w, "using HandleFunc")
	} else {
		io.WriteString(w, "using HandlerFunc")
	}
}

func d(w http.ResponseWriter, req *http.Request) {
	host := strings.Split(req.Host, ":")

	fmt.Fprintln(w, "in /dog route")

	if host[1] == "8080" {
		io.WriteString(w, "using HandleFunc")
	} else {
		io.WriteString(w, "using HandlerFunc")
	}
}

func useHandleFunc() {
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	// https://golang.org/pkg/net/http/#HandleFunc
	// HandlerFunc takes a string and a function with specific params as parameters
	http.HandleFunc("/cat", c)
	http.HandleFunc("/dog", d)

	http.ListenAndServe(":8080", nil)
}

func useHandlerFunc() {
	// func Handle(pattern string, handler Handler)
	// https://golang.org/pkg/net/http/#Handle
	// the Handle function takes a string, and a Handler as params

	// type HandlerFunc func(ResponseWriter, *Request)
	// https://golang.org/pkg/net/http/#HandlerFunc
	// Doc:
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

	// This means:
	// if a function f's parameters are defined the same as type HandlerFunc
	// which is: func(ResponseWriter, *Request)
	// HandlerFunc(f) is actually turning f into a Handler
	// and when this Handler runs, it is just calling f
	// see: https://golang.org/pkg/net/http/#HandlerFunc.ServeHTTP

	http.Handle("/cat", http.HandlerFunc(c))
	http.Handle("/dog", http.HandlerFunc(d))

	http.ListenAndServe(":8081", nil)
}

func main() {
	// comment and uncomment either function to see the difference
	// note that httpServeMux is listening on port 8080
	// while DefaultServerMux is listening on port 8081
	useHandleFunc()

	// useHandlerFunc()
}
