package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CmdTest(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "⚡ البوت شغال ويرد فوراً")
	bot.Send(reply)
}
