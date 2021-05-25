package observer

import "fmt"
// 通过观察者模式实现一个事件监听器
type Listener interface {
	Notify()
}

// OrderCreateListener 订单创建
type OrderCreateListener struct {}
func (listener OrderCreateListener) Notify() {
	fmt.Println("create order")
}

type StockDeductListener struct {}
func (listener StockDeductListener) Notify() {
	fmt.Println("deduct stock")
}

// Event 事件类
type Event struct {
	name string
	listeners []Listener
}

func (event *Event) Listen(listener Listener)  {
	event.listeners = append(event.listeners, listener)
}
func (event *Event) Trigger() {
	for _, listener := range event.listeners {
		listener.Notify()
	}
}

func NewEvent(name string) *Event {
	return &Event{name: name}
}