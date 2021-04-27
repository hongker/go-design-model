package abstractfactory

import "fmt"

type GoogleEmployee struct {}

func (employee GoogleEmployee) Work() {
	fmt.Println("google employee working..")
}

type GoogleDepartment struct {
	name string
}

func (department GoogleDepartment) Name() string {
	return fmt.Sprintf("google:%s", department.name)
}

type GoogleFactory struct {

}

func (factory GoogleFactory) CreateEmployee() Employee {
	return &GoogleEmployee{}
}

func (factory GoogleFactory) CreateDepartment(name string) Department {
	return &GoogleDepartment{name: name}
}

