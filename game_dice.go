package commands

import (
    "fmt"
    "math/rand"
    "time"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    ".."
)

func GameDice(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
    rand.Seed(time.Now().UnixNano())
    dice := rand.Intn(6) + 1

    uid := msg.From.ID
    prize := 0
    if dice >= 5 {
        prize = 200
    } else if dice >= 3 {
        prize = 50
    } else {
        prize = -30
    }
    newGold, _ := UpdateUserGold(uid, prize)
    AddMessage(uid)

    text := fmt.Sprintf("🎲 رمية النرد: %d\n💰 الذهب المكتسب: %d\n✨ رصيدك الآن: %d", dice, prize, newGold)
    bot.Send(tgbotapi.NewMessage(msg.Chat.ID, text))
}
