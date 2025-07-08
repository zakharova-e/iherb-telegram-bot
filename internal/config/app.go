package config

import (
	"log"
	"os"
)

var TelegramBotToken string

func ConfigInit(){
	token := os.Getenv("TELEGRAM_API_TOKEN")
	if token == ""{
		log.Fatalf("cannot load telegram api token")
	}
	TelegramBotToken = token
}