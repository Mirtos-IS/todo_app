package models

import (
	"database/sql"
)

type TodoList struct {
    Uid int
    Name string
    User_uid sql.NullInt64
    Count int
}

func GetLists() ([]TodoList, error){
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, _ := db.Query(
        `SELECT todo_lists.*, COUNT(todo_items.uid)
        FROM todo_lists
            INNER JOIN todo_items
            ON todo_lists.uid = todo_items.todo_list_uid
        GROUP BY todo_lists.uid
        ORDER BY uid DESC`)
    defer rows.Close()

    var lists []TodoList

    for rows.Next() {
        var list TodoList
        err = rows.Scan(&list.Uid, &list.Name, &list.User_uid, &list.Count)
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
