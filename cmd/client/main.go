package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"fmt"

	"short-news-bot/internal/initializers"
	"short-news-bot/internal/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
 
func main() {
	user := models.User{}

	user.User_ID = "001"

	err := initializers.DB.Create(&user).Error
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	fmt.Println("Бот запущен! Нажмите ctrl + c для выхода")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}