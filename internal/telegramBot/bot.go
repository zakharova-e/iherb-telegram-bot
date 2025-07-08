package telegrambot

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zakharova-e/iherb-telegram-bot/internal/config"
)


func Run() {
	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
		log.Printf("received a message: %v",update)
        if update.Message == nil {
            continue
        }
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
        msg.ReplyToMessageID = update.Message.MessageID
        if _, err := bot.Send(msg); err != nil {
            log.Println(err)
        }
    }
}