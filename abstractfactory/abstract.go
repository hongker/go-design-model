package abstractfactory

type EmployeeFactory interface {
	CreateEmployee() Employee
}

type DepartmentFactory interface {
	CreateDepartment(name string) Department
}

type CompanyFactory interface {
	EmployeeFactory
	DepartmentFactory
}

type Employee interface {
	Work()
}
type Department interface {
	Name() string
}

