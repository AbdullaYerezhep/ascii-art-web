package handler

import (
	"html/template"
	"net/http"

	asciiart "project/pkg/ascii-art"
)

var templ *template.Template

func IndexGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, "Invalid URL", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorPage(w, "Invalid Post Request URL", http.StatusMethodNotAllowed)
		return
	}

	templ = template.Must(template.ParseFiles("./ui/template/index.html"))

	templ.Execute(w, nil)
}

func Asciiart(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		ErrorPage(w, "Invalid URL", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		ErrorPage(w, "Invalid Post Request", http.StatusMethodNotAllowed)
		return
	}

	templ = template.Must(template.ParseFiles("./ui/template/index.html"))

	if r.Form.Has("text") || r.Form.Has("style") {
		ErrorPage(w, "Invalid Form Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	style := r.FormValue("style")

	if style != "standard" && style != "shadow" && style != "thinkertoy" {
		ErrorPage(w, "Invalid Form Request", http.StatusBadRequest)
		return
	}

	res, err := asciiart.AsciiArt(text, style)
	if err != nil {
		ErrorPage(w, "Invalid template", http.StatusBadRequest)
		return
	}
	templ.ExecuteTemplate(w, "index.html", res)
}
