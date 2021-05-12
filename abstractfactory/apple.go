package abstractfactory

import "fmt"


type AppleEmployee struct {}
func (employee AppleEmployee) Work() {
	fmt.Println("apple employee working..")
}

type AppleDepartment struct {
	name string
}
func (department AppleDepartment) Name() string {
	return fmt.Sprintf("apple:%s", department.name)
}

type AppleFactory struct {}
func (factory AppleFactory) CreateEmployee() Employee {
	return &AppleEmployee{}
}

func (factory AppleFactory) CreateDepartment(name string) Department {
	return &AppleDepartment{name: name}
}
