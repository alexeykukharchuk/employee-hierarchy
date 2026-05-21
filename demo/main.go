package main

import (
    "fmt"
    "example.com/employee"
)

func printEmployeeInfo(emp employee.Employee) {
    fmt.Printf("Должность: %s\n", emp.GetRole())
    fmt.Printf("Зарплата: %.2f руб.\n", emp.GetSalary())
    fmt.Printf("Работа: %s\n", emp.Work())
    fmt.Println("---")
}

func main() {
    // Создание TeamLead
    teamLead := employee.NewTeamLead("Алиса Иванова", 35, 95000, 8)
    
    // Создание Developer
    developer := employee.NewDeveloper("Пётр Смирнов", 28, 75000, "Go")
    
    // Создание Intern
    intern := employee.NewIntern("Дмитрий Козлов", 22, 35000, "Пётр Смирнов")
    
    fmt.Println("=== Демонстрация иерархии сотрудников ===\n")
    
    fmt.Println("Информация о тимлиде:")
    fmt.Printf("Имя: %s, Возраст: %d\n", teamLead.Name, teamLead.Age)
    fmt.Printf("Размер команды: %d человек\n", teamLead.TeamSize)
    fmt.Println(teamLead.Work())
    fmt.Println(teamLead.ConductMeeting())
    fmt.Printf("Зарплата: %.2f руб.\n\n", teamLead.GetSalary())
    
    fmt.Println("Информация о разработчике:")
    fmt.Printf("Имя: %s, Возраст: %d\n", developer.Name, developer.Age)
    fmt.Printf("Язык программирования: %s\n", developer.ProgrammingLanguage)
    fmt.Println(developer.Work())
    fmt.Println(developer.CodeReview())
    fmt.Printf("Зарплата: %.2f руб.\n\n", developer.GetSalary())
    
    fmt.Println("Информация о стажёре:")
    fmt.Printf("Имя: %s, Возраст: %d\n", intern.Name, intern.Age)
    fmt.Printf("Наставник: %s\n", intern.MentorName)
    fmt.Println(intern.Work())
    fmt.Println(intern.AskQuestion())
    fmt.Printf("Зарплата: %.2f руб.\n\n", intern.GetSalary())
    
    fmt.Println("=== Полиморфное поведение ===")
    employees := []employee.Employee{teamLead, developer, intern}
    for _, emp := range employees {
        printEmployeeInfo(emp)
    }
}
