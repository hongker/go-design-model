package factory

import "fmt"

type Logger interface {
	Write()
}

type FileLogger struct {}
func (logger FileLogger) Write() {
	fmt.Println("write to file...")
}

type DBLogger struct {}

func (logger DBLogger) Write() {
	fmt.Println("write to database...")
}

type Factory interface {
	Create() Logger
}

type FileLoggerFactory struct {}
func (factory FileLoggerFactory) Create() Logger{
	return new(FileLogger)
}

type DBLoggerFactory struct {}
func (factory DBLoggerFactory) Create() Logger {
	return new(DBLogger)
}


