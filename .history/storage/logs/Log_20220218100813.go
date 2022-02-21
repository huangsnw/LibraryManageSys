package logs

import "go.uber.org/zap"

var Log *zap.Logger

func Log() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	return logger
}
