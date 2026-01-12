package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
    var err error
    db, err = sql.Open("sqlite3", "./bot.db")
    if err != nil {
        return err
    }

    // إنشاء جدول المستخدمين إذا ما موجود
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            uid INTEGER PRIMARY KEY,
            name TEXT,
            gold INTEGER DEFAULT 0,
            total_messages INTEGER DEFAULT 0
        )
    `)
    return err
}

func GetUser(uid int64) (name string, gold int, totalMessages int, err error) {
    row := db.QueryRow("SELECT name, gold, total_messages FROM users WHERE uid = ?", uid)
    err = row.Scan(&name, &gold, &totalMessages)
    if err == sql.ErrNoRows {
        // إضافة مستخدم جديد تلقائي
        _, err = db.Exec("INSERT INTO users(uid, name) VALUES(?, ?)", uid, "")
        if err != nil {
            return
        }
        return "", 0, 0, nil
    }
    return
}

func UpdateUserGold(uid int64, amount int) (newGold int, err error) {
    _, err = db.Exec("UPDATE users SET gold = gold + ? WHERE uid = ?", amount, uid)
    if err != nil {
        return
    }
    row := db.QueryRow("SELECT gold FROM users WHERE uid = ?", uid)
    err = row.Scan(&newGold)
    return
}
