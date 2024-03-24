package main

import (
	"log"
	"net/http"
	"regexp"
	"todo/views"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func todoHandler(w http.ResponseWriter, r *http.Request) {
    views.Todo().Render(r.Context(), w)
}

func main() {
    http.HandleFunc("/todo/", todoHandler)

    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
