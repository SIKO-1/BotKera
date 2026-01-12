package commands

import (
    "fmt"
    "math/rand"
    "time"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDice(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    if update.Message.Text != "نرد" { // بدون شخطه
        return
    }

    rand.Seed(time.Now().UnixNano())
    value := rand.Intn(6) + 1

    msg := fmt.Sprintf("🎲 رميت النرد وطلعت: %d", value)
    bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
}
