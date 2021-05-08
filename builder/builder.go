package builder

import "fmt"

// Person 初始化该类时，有两个属性必填，两个属性选填
type Person struct {
	name string // 必填
	sex string // 必填
	age int // 选填
	birthday string // 选填
}

func (person Person) Info()  {
	fmt.Printf("名字:%s, 性别：%s,年龄：%d,生日：%s\n", person.name, person.sex, person.age, person.birthday)
}

// PersonBuilder Person的Builder接口，其中主要为Person的选填属性的setter方法
type PersonBuilder interface {
	SetAge(age int) *Builder
	SetBirthday(birthday string) *Builder
	Build() Person
}

type Builder struct {
	name string
	sex string
	age int
	birthday string
}

func (builder *Builder) SetAge(age int) *Builder {
	builder.age = age
	return builder
}

func (builder *Builder) SetBirthday(birthday string) *Builder {
	builder.birthday = birthday
	return builder
}

// NewBuilder 构造一个builder
func NewBuilder(name, sex string) PersonBuilder {
	return &Builder{name: name, sex: sex}
}

// Build 生成Person实例
func (builder *Builder) Build() Person {
	return Person{
		name:     builder.name,
		sex:      builder.sex,
		age:      builder.age,
		birthday: builder.birthday,
	}
}
