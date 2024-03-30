package models

import (
	"database/sql"
)

type TodoItem struct {
    Uid int
    Title string
    IsMarked bool
    TodoListUid sql.NullInt64
}

func GetItems(todoListUid int) ([]TodoItem, error) {
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, _ := db.Query("SELECT * FROM todo_items WHERE todo_list_uid = (?) ORDER BY uid DESC", todoListUid)
    defer rows.Close()

    var items []TodoItem

    for rows.Next() {
        var item TodoItem
        err = rows.Scan(&item.Uid, &item.Title, &item.TodoListUid, &item.IsMarked)
        if err != nil {
            panic(err)
        }
        items = append(items, item)
    }
    return items, nil
}


func SaveItem(title string, todoListUid int) (bool, error) {
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO todo_items(title, todo_list_uid) VALUES(?, ?)", title, todoListUid)
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

func ToggleComplete(uid int) (bool, error) {
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    _, err = db.Exec("UPDATE todo_items SET is_marked = NOT is_marked WHERE uid = (?)", uid)
    if err != nil {
        panic(err)
    }

    return true, nil
}
