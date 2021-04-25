package proxy

import "testing"

func TestNewSubject(t *testing.T) {
	subject := NewSubject()
	subject.Do()
}