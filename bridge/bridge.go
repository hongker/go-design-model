package bridge

import "fmt"

type Message interface {
	Send(content string, receiver string)
}

type Sms struct {}
func (Sms) Send(content string, receiver string) {
	fmt.Printf("send sms to %s: %s\n", receiver, content)
}



type Email struct {}
func (Email) Send(content string, receiver string) {
	fmt.Printf("send email to %s: %s\n", receiver, content)
}

type Notification interface {
	Notify(text, to string)
}

type AlertNotification struct {
	message Message
}
func (notification AlertNotification) Notify(text, to string) {
	notification.message.Send(fmt.Sprintf("[alert]%s", text), to)
}
func NewAlertNotification(message Message) Notification {
	return &AlertNotification{message: message}
}
type ErrorNotification struct {
	message Message
}
func (notification ErrorNotification) Notify(text, to string) {
	notification.message.Send(fmt.Sprintf("[error]%s", text), to)
}
func NewErrorNotification(message Message) Notification {
	return &ErrorNotification{message: message}
}

type Client struct {}

func (Client) Run() {
	alertNotification := NewAlertNotification(new(Email))
	alertNotification.Notify("something wrong", "tom")

	errorNotification := NewErrorNotification(new(Sms))
	errorNotification.Notify("something wrong", "bob")
}