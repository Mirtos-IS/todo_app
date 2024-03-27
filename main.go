package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"todo/models"
	"todo/views"

	_ "github.com/mattn/go-sqlite3"
)
var count int

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func todoHandler(w http.ResponseWriter, r *http.Request) {
    todoItems, _ := models.GetItems()
    views.Todo(todoItems).Render(r.Context(), w)
}

func todoNewHandler(w http.ResponseWriter, r *http.Request) {
    views.NewTodo().Render(r.Context(), w)
}

func todoAddHandler(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("title")
    models.SaveItem(title)
}

func todoDeleteHandler(w http.ResponseWriter, r *http.Request) {
    rawId := r.URL.Path[len("/todo/delete/"):]
    uid, _ := strconv.Atoi(rawId)
    models.DeleteItem(uid)
}

func main() {
    http.HandleFunc("/todo/", todoHandler)
    http.HandleFunc("/new/todo/", todoNewHandler)
    http.HandleFunc("/add/todo/", todoAddHandler)
    http.HandleFunc("/todo/delete/", todoDeleteHandler)
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/css"))))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
