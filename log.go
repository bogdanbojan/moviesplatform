package main

import (
	"log"
	"os"
)

type ServiceLogger struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func NewServiceLogger() *ServiceLogger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &ServiceLogger{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
}
