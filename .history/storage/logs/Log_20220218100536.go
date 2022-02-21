package logs

import "go.uber.org/zap"

func InitLog() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
}
