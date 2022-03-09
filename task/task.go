package task

import (
	"log"

	"github.com/robfig/cron/v3"
)

func PrintTime() {
	c := cron.New()
	c.AddFunc("@every 10s", tsk)
	c.Start()
}

func tsk() {
	log.Println("定时任务执行")
}
