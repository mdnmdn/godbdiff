package utils

import (
	"log"

	"go.uber.org/zap"
)

func Init() {

}

func GetLogger(name string) *zap.SugaredLogger {
	logger, err := zap.NewDevelopment()

	if err != nil {
		log.Fatal("err on zip init", err)
	}

	return logger.Sugar()
}
