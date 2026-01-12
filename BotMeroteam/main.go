// main.go
package main

import (
    "log"
    "os"
    "path/filepath"
    "plugin"
    "strings"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    botToken := os.Getenv("BOT_TOKEN")
    if botToken == "" {
        log.Fatal("❌ BOT_TOKEN غير معرف في البيئة")
    }

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
    log.Printf("🤖 البوت جاهز: %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    // ======================
    // تحميل جميع أوامر commands/
    // ======================
    commandHandlers := make(map[string]func(*tgbotapi.BotAPI, tgbotapi.Update))

    err = filepath.Walk("./commands", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && strings.HasSuffix(info.Name(), ".so") {
            plug, err := plugin.Open(path)
            if err != nil {
                log.Println("❌ خطأ بفتح البلجن:", path, err)
                return nil
            }
            sym, err := plug.Lookup("Handle")
            if err != nil {
                log.Println("❌ البلجن لا يحتوي Handle():", path)
                return nil
            }
            handler, ok := sym.(func(*tgbotapi.BotAPI, tgbotapi.Update))
            if !ok {
                log.Println("❌ Handle() مش من النوع الصحيح:", path)
                return nil
            }
            // اسم الأمر نفس اسم الملف بدون .so
            cmdName := strings.TrimSuffix(info.Name(), ".so")
            commandHandlers[cmdName] = handler
            log.Println("✅ تم تحميل الأمر:", cmdName)
        }
        return nil
    })
    if err != nil {
        log.Println("❌ خطأ بقراءة مجلد commands:", err)
    }

    log.Println("📦 جميع الأوامر جاهزة للعمل!")

    // ======================
    // الاستماع للتحديثات
    // ======================
    for update := range updates {
        if update.Message == nil {
            continue
        }

        text := strings.TrimSpace(update.Message.Text)
        if text == "" {
            continue
        }

        // تمرير الرسالة لأي أمر يطابق اسمها
        for name, handler := range commandHandlers {
            if strings.EqualFold(text, name) { // ما يحتاج "/"، فقط الاسم
                go handler(bot, update)
                break
            }
        }
    }
}
