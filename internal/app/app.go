package app

import (
	"net/http"
	"project/internal/handler"
)

func Run(port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexGet)
	mux.HandleFunc("/ascii-art", handler.Asciiart)
	styles := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", styles))
	return http.ListenAndServe(port, mux)
}
