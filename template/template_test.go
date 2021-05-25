package template

import "testing"

func TestDocSuper_Read(t *testing.T) {
	net := NewNetworkDoc()
	local := NewLocalDoc()

	net.Read()
	local.Read()
}