package flyweight

import (
	"fmt"
	"sync"
)

type Connection interface {
	Query()
}

type DBConnection struct {}

func (connection DBConnection) Query() {
	fmt.Printf("query from database\n")
}

type RedisConnection struct {}
func (connection RedisConnection) Query() {
	fmt.Printf("query from redis\n")
}

// Container 连接容器
type Container struct {
	mu sync.Mutex
	connections map[string]Connection
}

// GetConnection 根据类型获取连接
func (container *Container) GetConnection(name string) Connection {
	container.mu.Lock()
	defer container.mu.Unlock()
	conn, ok := container.connections[name]
	if ok {
		return conn
	}

	if name == "redis" {
		conn = new(RedisConnection)
	}else {
		conn = new(DBConnection)
	}
	container.connections[name] = conn
	return conn
}

func NewContainer() *Container {
	return &Container{
		connections: map[string]Connection{},
	}
}

