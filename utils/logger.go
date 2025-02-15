package utils

import "go.uber.org/zap"

var logger *zap.Logger

func InitLogger() *zap.Logger {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}

func GetLogger() *zap.Logger {
	return logger
}
