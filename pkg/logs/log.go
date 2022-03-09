package logs

import (
	"log"
)

func InitLog() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[系统日志]")
	// file, err := os.OpenFile("./log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatal("Log Config Error")
	// }
	// log.SetOutput(file)
}
