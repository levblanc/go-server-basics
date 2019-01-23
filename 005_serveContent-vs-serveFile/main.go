package main

import (
	"net/http"
	"os"
)

func geekCatHandler(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./assets/geekCat.png")

	if err != nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}

	defer file.Close()

	fileStat, err := file.Stat()

	if err != nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}

	http.ServeContent(w, req, file.Name(), fileStat.ModTime(), file)
}

func unicornCatHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./assets/unicornCat.jpg")
}

func main() {
	http.HandleFunc("/geekcat", geekCatHandler)
	http.HandleFunc("/unicorncat", unicornCatHandler)

	http.ListenAndServe(":8080", nil)
}
