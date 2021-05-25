package mediator

import "testing"

func TestApp(t *testing.T) {
	app := NewApp()
	app.InvokeUserHandler()
	app.InvokeGameHandler()
}