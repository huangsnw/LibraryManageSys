package logs

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func InitLog() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	Log = logger.Sugar()
}

func GetLog() *zap.SugaredLogger {
	return Log
}