package commands

import (
    "fmt"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    ".." // للوصول لـ db.go
)

func CmdStats(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
    uid := msg.From.ID
    name, gold, msgs, _ := GetUser(uid)

    text := fmt.Sprintf(`╔═════════════════╗
إحصائياتك
╚═════════════════╝

• الاسم: %s
• UID: %d
• الذهب: %d
• الرسائل: %d
`, name, uid, gold, msgs)

    bot.Send(tgbotapi.NewMessage(msg.Chat.ID, text))
}
