package main

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

// DB هو الاتصال بالقاعدة
var DB *sql.DB

// تهيئة قاعدة البيانات
func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./botmeroteam.db")
    if err != nil {
        log.Fatal(err)
    }

    // إنشاء جدول المستخدمين إذا ما كان موجود
    query := `
    CREATE TABLE IF NOT EXISTS users (
        uid INTEGER PRIMARY KEY,
        name TEXT,
        username TEXT,
        gold INTEGER DEFAULT 0,
        bank INTEGER DEFAULT 0,
        total_messages INTEGER DEFAULT 0,
        daily_usage INTEGER DEFAULT 0,
        banned BOOLEAN DEFAULT 0
    );
    `
    _, err = DB.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}

// ====================
// دوال مساعدة للمستخدمين
// ====================

// احصل على المستخدم حسب UID
func GetUser(uid int64) (map[string]interface{}, error) {
    row := DB.QueryRow("SELECT uid, name, username, gold, bank, total_messages, daily_usage, banned FROM users WHERE uid = ?", uid)

    var userID int64
    var name, username string
    var gold, bank, totalMessages, dailyUsage int
    var banned bool

    err := row.Scan(&userID, &name, &username, &gold, &bank, &totalMessages, &dailyUsage, &banned)
    if err == sql.ErrNoRows {
        // لو المستخدم غير موجود، نضيفه تلقائياً
        _, err := DB.Exec("INSERT INTO users(uid) VALUES(?)", uid)
        if err != nil {
            return nil, err
        }
        return GetUser(uid)
    } else if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "uid":            userID,
        "name":           name,
        "username":       username,
        "gold":           gold,
        "bank":           bank,
        "total_messages": totalMessages,
        "daily_usage":    dailyUsage,
        "banned":         banned,
    }, nil
}

// تحديث الذهب للمستخدم
func UpdateUserGold(uid int64, amount int) (int, error) {
    user, err := GetUser(uid)
    if err != nil {
        return 0, err
    }
    newGold := user["gold"].(int) + amount
    if newGold < 0 {
        newGold = 0
    }

    _, err = DB.Exec("UPDATE users SET gold = ? WHERE uid = ?", newGold, uid)
    if err != nil {
        return 0, err
    }
    return newGold, nil
}

// احصل على كل المستخدمين
func GetAllUsers() ([]map[string]interface{}, error) {
    rows, err := DB.Query("SELECT uid, name, username, gold, bank, total_messages, daily_usage, banned FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []map[string]interface{}
    for rows.Next() {
        var uid int64
        var name, username string
        var gold, bank, totalMessages, dailyUsage int
        var banned bool

        err := rows.Scan(&uid, &name, &username, &gold, &bank, &totalMessages, &dailyUsage, &banned)
        if err != nil {
            return nil, err
        }

        users = append(users, map[string]interface{}{
            "uid":            uid,
            "name":           name,
            "username":       username,
            "gold":           gold,
            "bank":           bank,
            "total_messages": totalMessages,
            "daily_usage":    dailyUsage,
            "banned":         banned,
        })
    }

    return users, nil
}
