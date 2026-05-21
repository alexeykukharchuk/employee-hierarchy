package employee

import "fmt"

type TeamLead struct {
    *BaseEmployee
    TeamSize int
}

func NewTeamLead(name string, age int, salary float64, teamSize int) *TeamLead {
    return &TeamLead{
        BaseEmployee: NewBaseEmployee(name, age, salary),
        TeamSize:     teamSize,
    }
}

func (t *TeamLead) GetRole() string {
    return "Тимлид"
}

func (t *TeamLead) Work() string {
    return fmt.Sprintf("%s управляет командой из %d разработчиков", t.Name, t.TeamSize)
}

func (t *TeamLead) ConductMeeting() string {
    return fmt.Sprintf("%s проводит ежедневное собрание", t.Name)
}
