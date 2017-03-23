package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[1:]
	file, err := os.OpenFile("data/"+filename+".html", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("unable to open a file")
	}
	mwriter := io.MultiWriter(file, w)

	// replace the macro in template
	t, _ := template.ParseFiles("tmpl/view.html")
	w.Header().Set("Content-Type", "text/html")

	p := &Page{Title: filename}
	t.Execute(mwriter, p)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
