package commands

import (
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleTestCommand يتحقق من رسالة المستخدم ويرد فوراً
func HandleTestCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    text := message.Text

    // إذا كتب المستخدم "اختبار" بدون شخطه
    if text == "اختبار" {
        reply := tgbotapi.NewMessage(message.Chat.ID, "✅ البوت يعمل تمام وسريع!")
        bot.Send(reply)
    }
}
