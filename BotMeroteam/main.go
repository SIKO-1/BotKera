package main

import (
    "log"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "BotMeroteam/commands"
)

func main() {
    botToken := os.Getenv("BOT_TOKEN")
    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil { // ignore non-Message updates
            continue
        }

        // استدعاء الأوامر بدون شخطه
        commands.HandleTest(bot, update)
        commands.HandleDice(bot, update)
        commands.HandleRoulette(bot, update)
    }
}
