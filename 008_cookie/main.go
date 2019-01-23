package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", renderIndex)
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/expire", expireCookie)
	http.ListenAndServe(":8080", nil)
}

func renderIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(
		w,
		`<h1><a href='/set'>Set Cookie</a></h1>`,
	)
}

func setCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit-count")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visit-count",
			Value: "0",
			// The Domain and Path directives define the scope of the cookie: what URLs the cookies should be sent to.
			// https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies#Scope_of_cookies
			Path: "/",
		}
	}

	visitCount, err := strconv.Atoi(cookie.Value)

	if err != nil {
		log.Fatalln(err)
	}

	visitCount++
	cookie.Value = strconv.Itoa(visitCount)

	http.SetCookie(w, cookie)
	fmt.Fprintln(
		w,
		`<h1>SET COOKIE SUCCESS!</h1> 
			<h1><a href="/read">Read Your Cookie</a></h1>`,
	)
}

func readCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit-count")

	if err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(
		w,
		`<h1>Your Cookie: %v</h1>
			<h1><a href="/set">Increment Cookie Count</a></h1>
			<h1><a href="/expire">Expire Cookie</a></h1>`,
		cookie,
	)
}

func expireCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit-count")

	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	cookie.MaxAge = -1

	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
