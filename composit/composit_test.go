package composit

import "testing"

func TestRouter(t *testing.T) {
	router := NewRouter("")
	router.AddNode(NewRouter("/hello"))

	user := NewRouter("/user")
	user.AddNode(NewRouter("/list"))
	user.AddNode(NewRouter("/create"))
	user.AddNode(NewRouter("/update"))
	user.AddNode(NewRouter("/delete"))
	router.AddNode(user)
	router.Handle("/user/list")
	router.Handle("/hello")
}
