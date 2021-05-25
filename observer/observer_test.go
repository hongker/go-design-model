package observer

import "testing"

func TestEvent_Trigger(t *testing.T) {
	event := NewEvent("OrderPayment")
	event.Listen(new(OrderCreateListener))
	event.Listen(new(StockDeductListener))

	event.Trigger()

}