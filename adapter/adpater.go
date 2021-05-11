package adapter

import "fmt"

// Target 客户端期待调用的接口
type Target interface {
	Request()
}

// Source 源类,未实现Target接口
type Source struct {}

// DoSomething 业务逻辑
func (Source) DoSomething() {
	fmt.Println("do something")
}

// Adapter 适配器，当客户端需要调用Target接口且Source类未能实现该接口
// Adapter就来代替Source类实现Target接口，达到适配的目的
type Adapter struct {
	Source
}

func (adapter Adapter) Request()  {
	adapter.DoSomething()
}