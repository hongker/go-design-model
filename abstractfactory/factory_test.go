package abstractfactory

import (
	"fmt"
	"testing"
)

func TestAppleFactory(t *testing.T)  {
	factory := new(AppleFactory)
	employee := factory.CreateEmployee()
	employee.Work()
	department := factory.CreateDepartment("test")
	fmt.Println(department.Name())
}


func TestGoogleFactory(t *testing.T)  {
	factory := new(GoogleFactory)
	employee := factory.CreateEmployee()
	employee.Work()
	department := factory.CreateDepartment("test")
	fmt.Println(department.Name())
}
