package task

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func PrintTime() {
	log.Println("Starting...")
	c := cron.New()
	c.AddFunc("0 */1 * * * ?", tsk)
	c.Start()
}

func tsk() {
	fmt.Println("定时任务！")
}
