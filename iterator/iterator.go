package iterator

type User struct {
	Name string
}

type UserRepository struct {
	items []User
}

func (repo *UserRepository) AddUser(user User) {
	repo.items = append(repo.items, user)
}

func (repo *UserRepository) GetIterator() func() (string, bool) {
	idx := 0
	return func() (name string, ok bool) {
		if idx >= len(repo.items) {
			return
		}
		name, ok = repo.items[idx].Name, true
		idx++
		return
	}
}

func NewUserRepository() *UserRepository {
	return &UserRepository{items: make([]User, 0, 10)}
}