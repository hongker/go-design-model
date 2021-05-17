package facade

import "testing"

func TestNewSyncFacade(t *testing.T) {
	facade := NewSyncFacade()
	facade.SyncGame()
	facade.SyncUser()
}