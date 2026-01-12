package commands

import (
    "math/rand"
    "time"

    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RouletteGame(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    if message.Text != "/روليت" {
        return
    }

    rand.Seed(time.Now().UnixNano())
    outcomes := []string{"💰 فوز!", "💀 خسارة!", "🎉 جاكبوت!"}
    result := outcomes[rand.Intn(len(outcomes))]

    msg := tgbotapi.NewMessage(message.Chat.ID, "🎰 روليت: "+result)
    bot.Send(msg)
}
