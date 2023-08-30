package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger(isDev bool) {
	var logger *zap.Logger
	var err error
	if isDev {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("failed to initialize logger, error: %v", err)
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("failed to initialize logger, error: %v", err)
		}
	}
	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
}
