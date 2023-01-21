package handler

import (
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, messege string, code int) {
	templ, err := template.ParseFiles("./ui/template/error.html")
	if err != nil {
		http.Error(w, "Invalid template", http.StatusInternalServerError)
		return
	}
	errorP := struct {
		Code    int
		Messege string
	}{
		Code:    code,
		Messege: messege,
	}
	templ.Execute(w, errorP)
}
