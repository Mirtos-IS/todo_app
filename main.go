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

func getUidFromUrl(r *http.Request, url string) int {
    rawId := r.URL.Path[len(url):]
    Uid, err := strconv.Atoi(rawId)
    if err != nil {
        panic(err)
    }
    return Uid
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
    todoListUid := getUidFromUrl(r, "/todo/")

    todoItems, _ := models.GetItems(todoListUid)
    views.Todo(todoItems, todoListUid).Render(r.Context(), w)
}

func todoNewHandler(w http.ResponseWriter, r *http.Request) {
    todoListUid := getUidFromUrl(r, "/new/todo/")

    views.NewTodo(todoListUid).Render(r.Context(), w)
}

func todoAddHandler(w http.ResponseWriter, r *http.Request) {
    todoListUid := getUidFromUrl(r, "/add/todo/")

    title := r.FormValue("title")
    w.Header().Set("HX-Refresh", "true")
    models.SaveItem(title, todoListUid)
}

func todoDeleteHandler(w http.ResponseWriter, r *http.Request) {
    uid := getUidFromUrl(r, "/todo/delete/")
    models.DeleteItem(uid)
}

func todoCheckCompleteHandler(w http.ResponseWriter, r *http.Request) {
    uid := getUidFromUrl(r, "/todo/check/toggle/")
    models.ToggleComplete(uid)
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
    //All Todo Lists Routes
    http.HandleFunc("/list/", listHandler)
    http.HandleFunc("/list/add", listAddHandler)
    http.HandleFunc("/list/create", listCreateHandler)
    //Single Todo List Routes
    http.HandleFunc("/todo/", todoHandler)
    http.HandleFunc("/new/todo/", todoNewHandler)
    http.HandleFunc("/add/todo/", todoAddHandler)
    http.HandleFunc("/todo/delete/", todoDeleteHandler)
    http.HandleFunc("/todo/check/toggle/", todoCheckCompleteHandler)
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/css"))))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
