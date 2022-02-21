package logs

import "go.uber.org/zap"

func GetLog() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	return logger
}
