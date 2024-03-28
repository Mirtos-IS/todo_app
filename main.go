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
    rawId := r.URL.Path[len("/todo/"):]
    todoListUid, _ := strconv.Atoi(rawId)

    todoItems, _ := models.GetItems(todoListUid)
    views.Todo(todoItems, todoListUid).Render(r.Context(), w)
}

func todoNewHandler(w http.ResponseWriter, r *http.Request) {
    rawId := r.URL.Path[len("/new/todo/"):]
    todoListUid, _ := strconv.Atoi(rawId)

    views.NewTodo(todoListUid).Render(r.Context(), w)
}

func todoAddHandler(w http.ResponseWriter, r *http.Request) {
    rawId := r.URL.Path[len("/add/todo/"):]
    todoListUid, _ := strconv.Atoi(rawId)

    title := r.FormValue("title")
    w.Header().Set("HX-Refresh", "true")
    models.SaveItem(title, todoListUid)
}

func todoDeleteHandler(w http.ResponseWriter, r *http.Request) {
    rawId := r.URL.Path[len("/todo/delete/"):]
    uid, _ := strconv.Atoi(rawId)
    models.DeleteItem(uid)
}

func todoCheckCompleteHandler(w http.ResponseWriter, r *http.Request) {
    rawId := r.URL.Path[len("/todo/check/complete/"):]
    uid, _ := strconv.Atoi(rawId)
    models.MarkAsComplete(uid)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    todoLists, _ := models.GetLists()
    views.TodoList(todoLists).Render(r.Context(), w)
}

func listAddHandler(w http.ResponseWriter, r *http.Request) {
    views.NewTodoList().Render(r.Context(), w)
}

func listCreateHandler(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")

    w.Header().Set("HX-Refresh", "true")
    models.CreateList(name)
}

func main() {
    http.HandleFunc("/list/", listHandler)
    http.HandleFunc("/list/add", listAddHandler)
    http.HandleFunc("/list/create", listCreateHandler)
    //Single Todo List Routes
    http.HandleFunc("/todo/", todoHandler)
    http.HandleFunc("/new/todo/", todoNewHandler)
    http.HandleFunc("/add/todo/", todoAddHandler)
    http.HandleFunc("/todo/delete/", todoDeleteHandler)
    http.HandleFunc("/todo/check/complete/", todoCheckCompleteHandler)
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/css"))))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
