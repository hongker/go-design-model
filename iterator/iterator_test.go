package iterator

import (
	"fmt"
	"testing"
)

func TestUserRepository_GetIterator(t *testing.T) {
	repo := NewUserRepository()
	repo.AddUser(User{Name: "bob"})
	repo.AddUser(User{Name: "alice"})
	repo.AddUser(User{Name: "jack"})
	repo.AddUser(User{Name: "somebody"})

	iterator := repo.GetIterator()
	for {
		name, ok := iterator()
		if !ok {
			break
		}
		fmt.Println(name)
	}
}
