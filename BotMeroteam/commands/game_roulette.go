package commands

import (
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GameRoulette(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	rand.Seed(time.Now().UnixNano())

	outcomes := []string{
		"🎉 فزت!",
		"💀 خسرت!",
		"🔥 جاكبوت!",
	}

	result := outcomes[rand.Intn(len(outcomes))]
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, result))
}
