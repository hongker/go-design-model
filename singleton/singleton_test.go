package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	instance := New(func() interface{} {
		return "hello"
	})
	assert.NotNil(t, instance)
}

func TestSingleton_Get(t *testing.T) {
	s := "hello"
	instance := New(func() interface{} {
		return s
	})
	assert.Equal(t, s, instance.Get())
}

func BenchmarkSingleton_Get(b *testing.B) {
	s := "hello"
	instance := New(func() interface{} {
		return s
	})
	for i := 0; i < b.N; i++ {
		instance.Get()
	}
}
