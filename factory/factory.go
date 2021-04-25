package factory

import "fmt"

type LoggerType uint

const(
	File LoggerType = iota + 1
	DB
)

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

type LoggerFactory struct {}
// 创建
func (factory LoggerFactory) Create(loggerType LoggerType) Logger{
	var logger Logger
	switch loggerType {
	case File:
		logger = new(FileLogger)
	case DB:
		logger = new(DBLogger)
	}
	return logger
}





