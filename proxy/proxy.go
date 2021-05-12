package proxy

import "fmt"
// Subject 抽象类
type Subject interface {
	Do()
}

// RealSubject 具体实现
type RealSubject struct {}
func (impl RealSubject) Do() {
	fmt.Println("do something")
}

// SubjectProxy 静态代理类
type SubjectProxy struct {
	subject RealSubject
}
func (proxy SubjectProxy) Do() {
	fmt.Println("proxy before")
	proxy.subject.Do()
	fmt.Println("proxy after")
}

func NewSubject() Subject {
	return new(SubjectProxy)
}
