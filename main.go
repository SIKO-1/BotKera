package main

import (
    "log"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "./commands"
)

func main() {
    token := os.Getenv("TELEGRAM_BOT_TOKEN")
    if token == "" {
        log.Fatal("Please set TELEGRAM_BOT_TOKEN environment variable")
    }

    if err := InitDB(); err != nil {
        log.Fatal(err)
    }

    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Fatal(err)
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil { 
            continue
        }

        go commands.Handle(bot, update.Message) // نرسل الرسالة لكل الأوامر
    }
}
