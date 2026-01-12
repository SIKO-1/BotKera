package main

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "✅ البوت شغال وسريع!")
    bot.Send(msg)
}
