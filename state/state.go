package state

import "fmt"

type State interface {
	DoAction(ctx *Context)
	ToString() string
}
type Context struct {
	state State
}

func (ctx *Context) State() State {
	return ctx.state
}

func NewContext() *Context {
	return &Context{}
}

func (ctx *Context) SetState(state State) {
	ctx.state = state
}

type StartState struct {}

func (s *StartState) DoAction(ctx *Context) {
	fmt.Println("start")
	ctx.SetState(s)
}

func (s *StartState) ToString() string {
	return "Start State"
}

type StopState struct {}
func (s *StopState) DoAction(ctx *Context) {
	fmt.Println("stop")
	ctx.SetState(s)
}

func (s *StopState) ToString() string {
	return "Stop State"
}
