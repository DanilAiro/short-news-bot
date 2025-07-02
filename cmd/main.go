package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"

	"short-news-bot/internal/initializers"
	"short-news-bot/internal/controllers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToTg()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
 
func main() {
	controllers.StartBotWork()
	
	log.Println("Бот запущен! Нажмите ctrl + c для выхода")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}