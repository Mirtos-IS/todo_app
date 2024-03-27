package models

import (
	"database/sql"
)

type TodoItem struct {
    Uid int
    Title string
}

func GetItems() ([]TodoItem, error) {
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, _ := db.Query("SELECT * FROM todo_items")
    defer rows.Close()

    var items []TodoItem

    for rows.Next() {
        var item TodoItem
        err = rows.Scan(&item.Uid, &item.Title)
        if err != nil {
            panic(err)
        }
        items = append(items, item)
    }
    return items, nil
}

func SaveItem(title string) (bool, error) {
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO todo_items(title) VALUES(?)", title)
    if err != nil {
        panic(err)
    }
    return true, nil
}

func DeleteItem(uid int) (bool, error) {
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM todo_items WHERE uid=(?)", uid)
    if err != nil {
        panic(err)
    }

    return true, nil
}
