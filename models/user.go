package models

import (
	"crypto/sha256"
	"todo/database"
)

type User struct {
    Uid int
    Username string
    password string
}

func createUser(username, password string) (*User, error) {
    db, err := database.OpenDB()
    if err != nil {
        return nil, err
    }

    defer db.Close()

    rows, err := db.Query("INSERT INTO users(username, password) Values(?,?)", username, password)
    if err != nil {
        return nil, err
    }
    var user *User
    for rows.Next() {
        rows.Scan(user.Uid, user.Username, user.password)
    }
    return user, nil
}

func UserUidIfExist(username, password string) (int, error) {
    db, err := database.OpenDB()
    if err != nil {
        return 0, nil
    }
    defer db.Close()

    rows, err := db.Query("SELECT uid FROM users WHERE username=(?) AND password=(?)", username, hashPassword(password))
    if err != nil {
        return 0, err
    }

    defer rows.Close()

    var user *User
    for rows.Next() {
        rows.Scan(user.Uid)
        return user.Uid, nil
    }

    return 0, nil
}

func getUser(uid int) (*User, error) {
    db, err := database.OpenDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()
    rows, err := db.Query("SELECT * FROM users WHERE uid=(?)", uid)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var user *User
    for rows.Next() {
        rows.Scan(user.Uid, user.Username, user.password)
    }
    return user, nil
}

func hashPassword(password string) (string) {
    hash := sha256.New()
    hash.Write([]byte(password))

    return string(hash.Sum(nil))
}
