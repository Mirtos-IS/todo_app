package models

import (
	"database/sql"
)

type TodoList struct {
    Uid int
    Name string
    User_uid sql.NullInt64
}

func GetLists() ([]TodoList, error){
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, _ := db.Query("SELECT * FROM todo_lists ORDER BY uid DESC")
    defer rows.Close()

    var lists []TodoList

    for rows.Next() {
        var list TodoList
        err = rows.Scan(&list.Uid, &list.Name, &list.User_uid)
        if err != nil {
            panic(err)
        }
        lists = append(lists, list)
    }
    return lists, nil
}

func CreateList(title string) (bool, error) {
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO todo_lists(name) VALUES(?)", title)
    if err != nil {
        panic(err)
    }
    return true, nil
}
