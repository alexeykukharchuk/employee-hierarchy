package employee

import "fmt"

type Developer struct {
    *BaseEmployee
    ProgrammingLanguage string
}

func NewDeveloper(name string, age int, salary float64, language string) *Developer {
    return &Developer{
        BaseEmployee:       NewBaseEmployee(name, age, salary),
        ProgrammingLanguage: language,
    }
}

func (d *Developer) GetRole() string {
    return "Разработчик"
}

func (d *Developer) Work() string {
    return fmt.Sprintf("%s пишет код на %s", d.Name, d.ProgrammingLanguage)
}

func (d *Developer) CodeReview() string {
    return fmt.Sprintf("%s проверяет пулл-реквесты", d.Name)
}
