package factory

import "testing"

func TestLoggerFactory_Create(t *testing.T) {
	factory := NewFileLoggerFactory()
	logger := factory.Create()
	logger.Write()
}