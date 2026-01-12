package main

import (
    "math/rand"
    "time"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    rand.Seed(time.Now().UnixNano())
    outcomes := []string{"فوز", "خسارة", "جاكبوت"}
    result := outcomes[rand.Intn(len(outcomes))]

    text := "🎰 روليت الإمبراطورية...\n"
    text += "النتيجة: " + result

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
    bot.Send(msg)
}
