package factory

import "testing"

func TestNewLoggerSimpleFactory(t *testing.T) {
	factory := NewLoggerSimpleFactory()
	logger := factory.Create(File)
	logger.Write()
}