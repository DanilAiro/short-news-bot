package controllers

import (
	"log"
	"context"
	"strconv"

	"short-news-bot/internal/initializers"
	"short-news-bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBotWork() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)

	// `updates` is a golang channel which receives telegram updates
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
			break
	}
}

func handleMessage(message *tgbotapi.Message) {
	user_chat_id := message.Chat.ID

	user := models.User{}

	user.User_ID = strconv.FormatInt(user_chat_id, 10)

	err := initializers.DB.Create(&user).Error
	if err != nil {
		log.Fatal("Error while creating new user")
	}

	msg := tgbotapi.NewMessage(user_chat_id, "Добавили вас в список, ожидайте новостей")

	_, err = initializers.BOT.Send(msg)
	if err != nil {
		log.Fatal("Error while sending message")
	}
}
