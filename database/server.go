package database

import "database/sql"

func OpenDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "database/todoApp.db")
    return db, err
}
