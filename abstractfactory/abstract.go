package abstractfactory

type CompanyFactory interface {
	CreateEmployee() Employee
	CreateDepartment(name string) Department
}

type Employee interface {
	Work()
}
type Department interface {
	Name() string
}

