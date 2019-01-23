package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit-count")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visit-count",
			Value: "0",
			Path:  "/",
		}
	}

	visitCount, err := strconv.Atoi(cookie.Value)

	if err != nil {
		log.Fatalln(err)
	}

	visitCount++
	cookie.Value = strconv.Itoa(visitCount)

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "SET COOKIE SUCCESS! Visit count is: ", cookie.Value)
}

func readCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit-count")

	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "COOKIE: ", cookie)
}
