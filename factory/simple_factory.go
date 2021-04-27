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
// LoggerSampleFactory 简单工厂类
type LoggerSimpleFactory struct {}
// Create 创建
func (factory LoggerSimpleFactory) Create(loggerType LoggerType) Logger{
	var logger Logger
	switch loggerType {
	case File:
		logger = new(FileLogger)
	case DB:
		logger = new(DBLogger)
	}
	return logger
}

func NewLoggerSimpleFactory() *LoggerSimpleFactory {
	return new(LoggerSimpleFactory)
}





