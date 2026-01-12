package main

import (
    "log"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "BotMeroteam/commands" // تأكد أن المجلد commands بنفس مسار main.go
)

func main() {
    botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
    if botToken == "" {
        log.Fatal("❌ ضع التوكن في متغير البيئة TELEGRAM_BOT_TOKEN")
    }

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic(err)
    }

    log.Printf("✅ بوت شغال باسم: %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        // هنا نستدعي أمر الاختبار
        commands.HandleTestCommand(bot, update.Message)
    }
}
