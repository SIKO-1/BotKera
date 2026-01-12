package commands

import (
    "math/rand"
    "time"

    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DiceGame(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    if message.Text != "/نرد" {
        return
    }

    rand.Seed(time.Now().UnixNano())
    value := rand.Intn(6) + 1
    msg := tgbotapi.NewMessage(message.Chat.ID, "🎲 رمية النرد: "+string(rune(value+'0')))
    bot.Send(msg)
}
