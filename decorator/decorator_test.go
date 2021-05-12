package decorator

import (
	"fmt"
	"testing"
)

func TestNewDecorator(t *testing.T) {
	component := &ConcreteComponent{}
	component.DoSomething()

	fmt.Println("decorator:")
	decorator := NewDecorator(component)
	decorator.DoSomething()

}