package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"todo/chat"
	"todo/gameOfLife"
	"todo/login"
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

func gameOfLifeHandler(w http.ResponseWriter, r *http.Request) {
    views.GameOfLife().Render(r.Context(), w)
}

func gameOfLifeGridHandler(w http.ResponseWriter, r *http.Request) {
    board := gameOfLife.NextBoardState()
    views.Grid(board).Render(r.Context(), w)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    views.Login().Render(r.Context(), w)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    views.Register().Render(r.Context(), w)
}

func registerCreateHandler(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")

    uid, err := models.CreateUser(username, password)
    if err != nil {
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte("Something went wrong"))
        w.Header().Set("HX-Refresh", "true")
        return
    }

    login.SetSession(uid, w)

    w.Header().Set("HX-Refresh", "true")
    http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func loginCheckHandler(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    uid, err := models.UserUidIfExist(username, password)
    if err != nil {
        panic(err)
    }

    if uid > 0 {
        login.SetSession(uid, w)
        http.Redirect(w, r, "/list", http.StatusFound)
        return
    }
    http.Redirect(w, r, "/login/", http.StatusFound)
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
    uid := login.GetUserId(r)
    if uid > 0 {
        todoLists, _ := models.GetLists(uid)
        views.TodoList(todoLists).Render(r.Context(), w)
        return
    }
    http.Redirect(w, r, "/login/", http.StatusFound)
}

func listAddHandler(w http.ResponseWriter, r *http.Request) {
    views.NewTodoList().Render(r.Context(), w)
}

func listCreateHandler(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    uid := login.GetUserId(r)

    w.Header().Set("HX-Refresh", "true")
    models.CreateList(name, uid)
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
    views.Chat().Render(r.Context(), w)
}

func main() {
    //login Routes
    http.HandleFunc("/login/", loginHandler)
    http.HandleFunc("/register/", registerHandler)
    http.HandleFunc("/register/create", registerCreateHandler)
    http.HandleFunc("/login/check", loginCheckHandler)
    //All Todo Lists Routes
    http.HandleFunc("/list/", listHandler)
    http.HandleFunc("/list/add", listAddHandler)
    http.HandleFunc("/list/create", listCreateHandler)
    //easter egg
    http.HandleFunc("/game-of-life/", gameOfLifeHandler)
    http.HandleFunc("/game-of-life/grid", gameOfLifeGridHandler)
    //Single Todo List Routes
    http.HandleFunc("/todo/", todoHandler)
    http.HandleFunc("/new/todo/", todoNewHandler)
    http.HandleFunc("/add/todo/", todoAddHandler)
    http.HandleFunc("/todo/delete/", todoDeleteHandler)
    http.HandleFunc("/todo/check/toggle/", todoCheckCompleteHandler)
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/css"))))

    //chat
    http.HandleFunc("/chat", ChatHandler)
    chat.StartServer()

    log.Fatal(http.ListenAndServe(":8080", nil))
    //chat server
    log.Fatal(http.ListenAndServe(":3000", nil))
}
