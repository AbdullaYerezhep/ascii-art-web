package app

import (
	"net/http"
	"project/internal/handler"
)

func Run(port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Index)
	// mux.HandleFunc("/showsmth", handler.Showsmth)
	styles := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", styles))
	return http.ListenAndServe(port, mux)
}
