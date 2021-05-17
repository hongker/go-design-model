package facade

import "fmt"

// Synchronous 同步模块
type Synchronous interface {
	Sync()
}

// Game 游戏类
type Game struct {}
func (game Game) Sync() {
	fmt.Println("sync game info")
}
// User 用户类
type User struct {}
func (user User) Sync() {
	fmt.Println("sync user info")
}

// SyncFacade 包含游戏和用户的同步
type SyncFacade struct {
	game Game
	user User
}

func (facade SyncFacade) SyncGame() {
	facade.game.Sync()
}
func (facade SyncFacade) SyncUser() {
	facade.user.Sync()
}

func NewSyncFacade() SyncFacade {
	return SyncFacade{
		game: Game{},
		user: User{},
	}
}

