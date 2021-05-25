package state

import (
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := NewContext()
	startState := new(StartState)
	startState.DoAction(ctx)

	stopState := new(StopState)
	stopState.DoAction(ctx)
	fmt.Printf("current state:%s\n", ctx.State().ToString())
}