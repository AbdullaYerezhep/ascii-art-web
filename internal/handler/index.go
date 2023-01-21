package handler

import (
	"fmt"
	"html/template"
	"net/http"

	asciiart "project/pkg/ascii-art"
)

var templ *template.Template

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, "Invalid URL", http.StatusNotFound)
		return
	}

	templ = template.Must(template.ParseFiles("./ui/template/index.html"))

	switch r.Method {
	case http.MethodGet:
		indexGet(w, r)
	case http.MethodPost:
		indexPost(w, r)
	default:
		ErrorPage(w, "Не правильный метод", http.StatusMethodNotAllowed)
	}
}

func indexPost(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	style := r.FormValue("style")

	res, err := asciiart.AsciiArt(text, style)
	fmt.Println(res)
	if err != nil {
		ErrorPage(w, "Invalid template", http.StatusInternalServerError)
		return
	}
	templ.Execute(w, res)
}

func indexGet(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, nil)
}
