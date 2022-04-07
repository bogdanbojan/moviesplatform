package api

import (
	"log"
	"os"
)

// ServiceLogger represents a centralized error and info log for our api.
type ServiceLogger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

// NewServiceLogger is the factory function for creating a ServiceLogger.
func NewServiceLogger() *ServiceLogger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &ServiceLogger{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}
}
