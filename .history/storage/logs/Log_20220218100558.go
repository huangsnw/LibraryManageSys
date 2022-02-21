package logs

import "go.uber.org/zap"

func InitLog() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
}
