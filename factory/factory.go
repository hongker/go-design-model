package factory


type LoggerFactory interface {
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

func NewFileLoggerFactory() LoggerFactory {
	return new(FileLoggerFactory)
}

func NewDBLoggerFactory() LoggerFactory {
	return new(FileLoggerFactory)
}
