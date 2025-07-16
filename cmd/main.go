package main

import (
	"os"
	"os/signal"
	"syscall"

	"short-news-bot/internal/controllers"
	"short-news-bot/internal/initializers"
)

func init() {
	initializers.ConnectToLogger()
	initializers.LoadEnvVariables()
	initializers.ConnectToTg()
	initializers.ConnectToCurApi()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	initializers.ConnectToCron()
}

func main() {
	controllers.UpdateCurrencies()
	controllers.StartBotWork()

	// Каждый день в 22:00
    initializers.AddCronJob("0 22 * * *", controllers.EveryDaySender)

	initializers.Log.Println("Бот запущен! Нажмите ctrl + c для выхода")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}