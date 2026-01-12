package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"BotMeroteam/commands"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("❌ BOT_TOKEN غير موجود")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false
	log.Println("🤖 البوت اشتغل")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		text := update.Message.Text

		switch text {
		case "تست":
			commands.CmdTest(bot, update.Message)

		case "نرد":
			commands.GameDice(bot, update.Message)

		case "روليت":
			commands.GameRoulette(bot, update.Message)
		}
	}
}
