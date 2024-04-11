package models

import (
	"database/sql"
	"todo/database"
)

type TodoList struct {
    Uid int
    Name string
    User_uid sql.NullInt64
    Count int
}

func GetLists(uid int) ([]TodoList, error){
    db, err := database.OpenDB()
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, _ := db.Query(
        `SELECT todo_lists.*, COUNT(todo_items.uid)
        FROM todo_lists
            FULL OUTER JOIN todo_items
            ON todo_lists.uid = todo_items.todo_list_uid
        WHERE user_uid = (?)
        GROUP BY todo_lists.uid
        ORDER BY uid DESC`, uid)
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

func CreateList(title string, uid int) (bool, error) {
    db, err := database.OpenDB()
    if err != nil {
        panic(err)
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO todo_lists(name, user_uid) VALUES(?, ?)", title, uid)
    if err != nil {
        panic(err)
    }
    return true, nil
}
