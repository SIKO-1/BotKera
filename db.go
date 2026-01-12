package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./bot.db")
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS users (
        uid INTEGER PRIMARY KEY,
        name TEXT,
        gold INTEGER DEFAULT 0,
        total_messages INTEGER DEFAULT 0
    );`

    _, err = DB.Exec(createTable)
    if err != nil {
        log.Fatal(err)
    }
}

// ==================
// وظائف بسيطة للبيانات
// ==================
func GetUser(uid int64) (name string, gold int, msgs int, err error) {
    row := DB.QueryRow("SELECT name, gold, total_messages FROM users WHERE uid = ?", uid)
    err = row.Scan(&name, &gold, &msgs)
    if err == sql.ErrNoRows {
        _, err = DB.Exec("INSERT INTO users(uid, name) VALUES(?, ?)", uid, "")
        if err != nil {
            return
        }
        name, gold, msgs = "", 0, 0
        err = nil
    }
    return
}

func UpdateUserGold(uid int64, amount int) (newGold int, err error) {
    _, err = DB.Exec("INSERT OR IGNORE INTO users(uid) VALUES(?)", uid)
    if err != nil {
        return
    }
    _, err = DB.Exec("UPDATE users SET gold = gold + ? WHERE uid = ?", amount, uid)
    if err != nil {
        return
    }
    row := DB.QueryRow("SELECT gold FROM users WHERE uid = ?", uid)
    err = row.Scan(&newGold)
    return
}

func AddMessage(uid int64) error {
    _, err := DB.Exec("INSERT OR IGNORE INTO users(uid) VALUES(?)", uid)
    if err != nil {
        return err
    }
    _, err = DB.Exec("UPDATE users SET total_messages = total_messages + 1 WHERE uid = ?", uid)
    return err
}
