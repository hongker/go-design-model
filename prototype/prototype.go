package prototype

type Student struct {
	Name string
	Books []string
}

// Clone 拷贝原型对象
func (student *Student) Clone() *Student {
	// 浅拷贝:更改拷贝项中的指针或引用对象，源自同一个原型的全部的克隆对象都会受到影响
	return student
}

// DeepClone 深拷贝：重新申请内存空间，更改拷贝项中的指针或引用对象，不会影响其他克隆对象
func (student *Student) DeepClone() *Student {

	books := make([]string, 0, len(student.Books))
	for _, book := range student.Books {
		books = append(books, book)
	}
	return &Student{
		Name:  student.Name,
		Books: student.Books,
	}
}

