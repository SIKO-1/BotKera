package main

import (
    "log"
    "strconv"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "./commands"
)

func main() {
    bot, err := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")
    if err != nil {
        log.Fatal(err)
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    InitDB()

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        text := update.Message.Text

        // أوامر بسيطة
        switch {
        case text == "احصائيات":
            commands.CmdStats(bot, update.Message)
        case text == "نرد":
            commands.GameDice(bot, update.Message)
        case len(text) > 6 && text[:6] == "روليت ":
            bet, err := strconv.Atoi(text[6:])
            if err == nil {
                commands.GameRoulette(bot, update.Message, bet)
            } else {
                bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "❌ يجب كتابة رقم صحيح بعد 'روليت'"))
            }
        }
    }
}
