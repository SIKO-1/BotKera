package commands

import (
    "fmt"
    "math/rand"
    "time"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleRoulette(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    if update.Message.Text != "روليت" { // بدون شخطه
        return
    }

    rand.Seed(time.Now().UnixNano())
    outcomes := []string{"خسارة", "فوز", "جاكبوت"}
    weights := []int{50, 45, 5} // احتمالات تقريبة

    total := 0
    for _, w := range weights {
        total += w
    }

    roll := rand.Intn(total)
    var result string
    sum := 0
    for i, w := range weights {
        sum += w
        if roll < sum {
            result = outcomes[i]
            break
        }
    }

    msg := fmt.Sprintf("🎰 نتيجة الروليت: %s", result)
    bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
}
