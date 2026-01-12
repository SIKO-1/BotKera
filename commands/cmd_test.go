package commands

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// أمر اختبار سريع
func TestCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    if message.Text == "/اختبار" {
        msg := tgbotapi.NewMessage(message.Chat.ID, "✅ البوت شغال ويرد بسرعة!")
        bot.Send(msg)
    }
}
