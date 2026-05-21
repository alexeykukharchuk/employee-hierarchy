package employee

import "fmt"

type Intern struct {
    *BaseEmployee
    MentorName string
}

func NewIntern(name string, age int, salary float64, mentorName string) *Intern {
    return &Intern{
        BaseEmployee: NewBaseEmployee(name, age, salary),
        MentorName:   mentorName,
    }
}

func (i *Intern) GetRole() string {
    return "Стажёр"
}

func (i *Intern) Work() string {
    return fmt.Sprintf("%s учится у наставника %s", i.Name, i.MentorName)
}

func (i *Intern) AskQuestion() string {
    return fmt.Sprintf("%s задаёт вопрос %s", i.Name, i.MentorName)
}
