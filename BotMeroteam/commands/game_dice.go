package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GameDice(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	dice := tgbotapi.NewDice(msg.Chat.ID)
	bot.Send(dice)
}
