package main

import (
    "math/rand"
    "time"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    rand.Seed(time.Now().UnixNano())
    value := rand.Intn(6) + 1

    text := "🎲 رمي النرد...\n"
    text += "الحظ يبتسم لك! الرقم: " + string(rune('0'+value))

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
    bot.Send(msg)
}
