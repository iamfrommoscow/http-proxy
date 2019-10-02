package helpers

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func LogMessage(message string) {
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info(
		zap.String("message", message),
	)
}
