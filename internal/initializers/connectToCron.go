package initializers

import (
	"github.com/robfig/cron/v3"
)

var CRON *cron.Cron

func ConnectToCron() {
	CRON = cron.New()

	CRON.Start()
}

func AddCronJob(schedule string, job func()) error {
	_, err := CRON.AddFunc(schedule, job)
	return err
}