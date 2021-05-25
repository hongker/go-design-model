package chain

import "testing"

func TestFilter(t *testing.T) {
	adFilter := new(AdFilter)
	sensitiveFilter := new(SensitiveFilter)
	adFilter.SetNext(sensitiveFilter)

	adFilter.Handle("some ad, somebody is sb")
}