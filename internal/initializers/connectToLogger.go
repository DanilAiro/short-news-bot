package initializers

import (
	"os"
	"log"
)

var (
	Log *log.Logger
)

func ConnectToLogger() {
	logFile, err := os.OpenFile("history.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}

	Log = log.New(logFile, "", log.LstdFlags)
}
