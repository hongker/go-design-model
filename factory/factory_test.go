package factory

import "testing"

func TestLoggerFactory_Create(t *testing.T) {
	factory := new(LoggerFactory)
	logger := factory.Create(File)
	logger.Write()
}