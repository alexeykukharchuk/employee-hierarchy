package employee

type Employee interface {
    GetRole() string
    GetSalary() float64
    Work() string
}

type BaseEmployee struct {
    Name   string
    Age    int
    Salary float64
}

func NewBaseEmployee(name string, age int, salary float64) *BaseEmployee {
    return &BaseEmployee{Name: name, Age: age, Salary: salary}
}

func (e *BaseEmployee) GetSalary() float64 {
    return e.Salary
}
