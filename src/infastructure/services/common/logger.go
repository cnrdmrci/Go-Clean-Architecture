package common

import (
	"go-clean-architecture/src/core/application/interfaces/services"
	"log"
)

type LoggerService struct{}

func CreateLogger() interface_services.LoggerService {
	return &LoggerService{}
}

func (l *LoggerService) LogError(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *LoggerService) LogAccess(format string, v ...interface{}) {
	log.Printf(format, v...)
}
