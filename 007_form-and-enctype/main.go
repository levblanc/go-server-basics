package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", renderIndex)
	http.HandleFunc("/url", renderURL)
	http.HandleFunc("/form", renderFormValue)
	http.HandleFunc("/file-upload", renderFileUpload)
	http.HandleFunc("/enctype", renderEnctype)

	http.ListenAndServe(":8080", nil)
}

func renderIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8;")
	tpl.ExecuteTemplate(w, "index.html", "")
}

func renderURL(w http.ResponseWriter, req *http.Request) {
	query := req.FormValue("q")

	w.Header().Set("Content-Type", "text/html; charset=utf8;")
	tpl.ExecuteTemplate(w, "url.html", query)
}

func renderFormValue(w http.ResponseWriter, req *http.Request) {
	query := req.FormValue("inputValue")

	data := struct {
		Method   string
		FormData string
	}{
		req.Method,
		query,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf8;")
	tpl.ExecuteTemplate(w, "form.html", data)
}

func renderEnctype(w http.ResponseWriter, req *http.Request) {
	// read request body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	// convert byte slice to string
	reqBody := string(bs)

	w.Header().Set("Content-Type", "text/html;charset=utf8;")
	tpl.ExecuteTemplate(w, "enctype.html", reqBody)
}

func renderFileUpload(w http.ResponseWriter, req *http.Request) {
	type data struct {
		FileContent   string
		UploadSuccess bool
	}

	var s string
	var d data

	if req.Method == http.MethodPost {
		file, header, err := req.FormFile("file")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		// read file contents
		bs, err := ioutil.ReadAll(file)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)

		// store file content into a file on server
		f, err := os.Create(filepath.Join("./file/", header.Filename))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()

		_, err = f.Write(bs)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		d = data{
			FileContent:   s,
			UploadSuccess: true,
		}
	}

	w.Header().Set("Content-Type", "text/html;charset=utf8;")
	tpl.ExecuteTemplate(w, "upload.html", d)
}
