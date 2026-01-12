package commands

import (
    "fmt"
    "math/rand"
    "time"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    ".."
)

func GameRoulette(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, bet int) {
    uid := msg.From.ID
    gold, _ := UpdateUserGold(uid, 0)
    if bet > gold {
        bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "💸 رصيدك غير كافي"))
        return
    }

    rand.Seed(time.Now().UnixNano())
    outcome := rand.Intn(100)
    prize := 0
    var result string

    switch {
    case outcome < 5:
        prize = bet * 5
        result = "جاكبوت أسطوري"
    case outcome < 50:
        prize = -bet
        result = "خسارة"
    default:
        prize = bet
        result = "فوز"
    }

    newGold, _ := UpdateUserGold(uid, prize)
    AddMessage(uid)

    text := fmt.Sprintf("🎰 روليت: %s\n💰 الذهب المكتسب: %d\n✨ رصيدك الآن: %d", result, prize, newGold)
    bot.Send(tgbotapi.NewMessage(msg.Chat.ID, text))
}
