package memento

import (
	"fmt"
	"testing"
)

func TestOriginator_RecoverByMemento(t *testing.T) {
	originator := NewOriginator()
	careTaker := NewCareTaker()
	originator.SetState("waiting")
	originator.SetState("doing")
	careTaker.Add(originator.GenerateMemento())
	originator.SetState("done")
	careTaker.Add(originator.GenerateMemento())
	originator.SetState("finished")

	fmt.Println("current state:", originator.GetState())
	originator.RecoverByMemento(careTaker.Get(0))
	fmt.Println("first memento state:", originator.GetState())
	originator.RecoverByMemento(careTaker.Get(1))
	fmt.Println("second memento state:", originator.GetState())


}