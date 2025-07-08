package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/zakharova-e/iherb-telegram-bot/internal/config"
	"github.com/zakharova-e/iherb-telegram-bot/internal/telegramBot"
)

func main() {
	if err := godotenv.Load(); err != nil {
		//cannot continue without env file, because of Telegram Api Token
		log.Fatal(err)
	}
	config.ConfigInit()
    telegrambot.Run()
}