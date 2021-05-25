package mediator

import "fmt"

// App 应用类通过网关调用userHandler和gameHandler两个接口
// 好处是如果被调用方接口有所变化，不必调整app逻辑，在gateway这个中介者类修改即可
type App struct {
	gateway Gateway
}

func NewApp() *App {
	return &App{gateway: NewGateway()}
}

func (app *App) InvokeUserHandler() {
	app.gateway.GetHandler("user").Handle()
}
func (app *App) InvokeGameHandler() {
	app.gateway.GetHandler("game").Handle()
}

type Gateway interface {
	GetHandler(name string) Handler
}

type gateway struct {
	handlers map[string]Handler
}

func (gw *gateway) GetHandler(name string) Handler{
	return gw.handlers[name]
}

func NewGateway() Gateway  {
	return &gateway{handlers: map[string]Handler{
		"user": new(UserHandler),
		"game": new(GameHandler),
	}}
}

type Handler interface {
	Handle()
}
type UserHandler struct {}

func (handler UserHandler) Handle() {
	fmt.Println("user handled")
}

type GameHandler struct {

}
func (handler GameHandler) Handle() {
	fmt.Println("game handled")
}