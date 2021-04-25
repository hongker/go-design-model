package singleton

import "sync"

// Singleton 单例模式
type Singleton struct {
	// once
	once sync.Once
	// object constructor
	constructor Constructor
	// user object
	instance interface{}
}

// Constructor 构造函数
type Constructor func() interface{}

// Instance return target object
func (singleton *Singleton) Instance() interface{} {
	singleton.once.Do(func() {
		// init once
		singleton.instance = singleton.constructor()
	})
	return singleton.instance
}

// New 实例化
func New(constructor Constructor) *Singleton {
	return &Singleton{constructor: constructor}
}
