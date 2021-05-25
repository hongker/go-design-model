package memento

// Memento 备忘录类
type Memento struct {
	state string
}

func NewMemento(state string) Memento {
	return Memento{state: state}
}

func (memento Memento) GetState() string {
	return memento.state
}

// Originator 原始类
type Originator struct {
	state string
}

func (originator *Originator) SetState(state string) {
	originator.state = state
}

func NewOriginator() *Originator {
	return &Originator{}
}


func (originator *Originator) GetState() string {
	return originator.state
}

func (originator *Originator) GenerateMemento() Memento {
	return NewMemento(originator.state)
}

func (originator *Originator) RecoverByMemento(memento Memento) {
	originator.state = memento.GetState()
}

// CareTaker 存储备忘录
type CareTaker struct {
	mementoList []Memento
}

func NewCareTaker() *CareTaker {
	return &CareTaker{mementoList: []Memento{}}
}
func (ct *CareTaker) Add(memento Memento) {
	ct.mementoList = append(ct.mementoList, memento)
}
func (ct *CareTaker) Get(idx int) Memento{
	if len(ct.mementoList) <= idx || idx<0 {
		panic("out of index")
	}
	return ct.mementoList[idx]
}

