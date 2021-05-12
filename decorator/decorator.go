package decorator

import "fmt"

type Component interface {
	DoSomething()
}

type ConcreteComponent struct {
	
}

func (component ConcreteComponent) DoSomething() {
	fmt.Println("do something")
}

type Decorator struct {
	component Component
}

func NewDecorator(component Component) *Decorator {
	return &Decorator{component: component}
}

func (decorator Decorator) DoSomething() {
	fmt.Println("do decorator")
	decorator.component.DoSomething()
}