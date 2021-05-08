package builder

import "testing"

func TestNewBuilder(t *testing.T) {
	person := NewBuilder("张三", "男").
		SetAge(26).
		SetBirthday("1994/01/01").
		Build()

	person.Info()
}