package commands

import (
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleTest(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    if update.Message.Text != "اختبار" { // بدون شخطه
        return
    }

    bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "✅ البوت شغال!"))
}
