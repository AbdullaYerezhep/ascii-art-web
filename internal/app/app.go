package app

import (
	"fmt"
	"net/http"
	"project/internal/handler"
)

func Run(port string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexGet)
	mux.HandleFunc("/ascii-art", handler.Asciiart)	
	styles := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", styles))
	img := http.FileServer(http.Dir("./ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets", img))
	fmt.Println("Server is up and running!")
	return http.ListenAndServe(port, mux)
}
