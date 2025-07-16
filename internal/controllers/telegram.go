package controllers

import (
	"context"
	"strconv"

	"short-news-bot/internal/initializers"
	"short-news-bot/internal/models"

	"gorm.io/gorm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBotWork() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ctx := context.Background()

	updates := initializers.BOT.GetUpdatesChan(u)

	// Pass cancellable context to goroutine
	go receiveUpdates(ctx, updates)
}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {
	switch {
		case update.Message != nil:
			handleMessage(update.Message)
	}
}

func handleMessage(message *tgbotapi.Message) {
	user_chat_id := message.Chat.ID
	user_id := user_chat_id

	user := models.User{}
	user.User_ID = user_id
	
	err := initializers.DB.Where("user_id = ?", user_id).First(&user).Error

	if err == nil {
		sendMessage(user_chat_id, "Вы уже есть в списке, ожидайте новостей")
		return
	}

	if err == gorm.ErrRecordNotFound {
		user := models.User{User_ID: user_id}
		err := initializers.DB.Create(&user).Error

		if err != nil {
			initializers.Log.Fatal("Error while creating new user: ", err.Error())
		}

		sendMessage(user_chat_id, "Добавили вас в список, ожидайте новостей")
	} else {
		initializers.Log.Fatal("Error while checking user existence: ", err.Error())
	}
}

func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := initializers.BOT.Send(msg)

	if err != nil {
		initializers.Log.Fatal("Error while sending message: ", err.Error())
	}
}

func EveryDaySender() {
	UpdateCurrencies()

	var text string

	if len(AllCurrencies) > 0 {
		text += "Курсы валют:\n"

		for k, v := range AllCurrencies {
			text += k + " - " + strconv.FormatFloat(v, 'f', 0, 64) + " RUB\n"
		}

		text += "\n"
	}

	users := []models.User{}
	initializers.DB.Find(&users)

	if len(users) > 0 {
		for _, v := range users {
			sendMessage(v.User_ID, text)
		}

		return
	}
}