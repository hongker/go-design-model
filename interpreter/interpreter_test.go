package interpreter

import (
	"fmt"
	"testing"
)

func TestNewOrExpression(t *testing.T) {
	lee := NewTerminalExpression("lee")
	wan := NewTerminalExpression("wan")
	isMale := NewOrExpression(lee, wan)
	fmt.Println("lee is male?", isMale.Interpreter("lee"))
}

func TestNewAndExpression(t *testing.T) {
	alice := NewTerminalExpression("alice")
	married := NewTerminalExpression("married")
	isMarried := NewAndExpression(alice, married)
	fmt.Println("alice is a married women?", isMarried.Interpreter("married alice"))
}