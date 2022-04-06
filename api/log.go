package api

import (
	"log"
	"os"
)

type ServiceLogger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func NewServiceLogger() *ServiceLogger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &ServiceLogger{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}
}
