package flyweight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer_GetConnection(t *testing.T) {
	container := NewContainer()
	redis := container.GetConnection("redis")
	redis.Query()

	db := container.GetConnection("db")
	db.Query()

	dbOther := container.GetConnection("db")
	assert.Equal(t, db, dbOther)

}