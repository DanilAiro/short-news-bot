package initializers

import (
	"os"
	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BOT *tgbotapi.BotAPI

func ConnectToTg() {
	var err error
	BOT, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	
	if err != nil {
		panic("Failed to connect to tg: " + err.Error())
	}
}