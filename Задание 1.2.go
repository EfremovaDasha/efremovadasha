package main

import "fmt"

type Employee struct {
    Name string
    Age  int
}

type OfficeManager struct {
    Employees [512]Employee
    Count     int
}

func main() {
    om := &OfficeManager{}

    // Добавление сотрудников
    om.AddEmployee("John Doe", 30)
    om.AddEmployee("Jane Smith", 35)
    om.AddEmployee("Bob Johnson", 40)

    // Вывод списка сотрудников
    om.PrintEmployees()

    // Удаление сотрудника
    om.RemoveEmployee("Jane Smith")

    // Вывод списка сотрудников после удаления
    om.PrintEmployees()
}

func (om *OfficeManager) AddEmployee(name string, age int) {
    if om.Count >= 512 {
        fmt.Println("Достигнут максимальный лимит сотрудников (512)")
        return
    }

    om.Employees[om.Count] = Employee{Name: name, Age: age}
    om.Count++
}

func (om *OfficeManager) RemoveEmployee(name string) {
    for i := 0; i < om.Count; i++ {
        if om.Employees[i].Name == name {
            // Сдвигаем все элементы после удаляемого
            for j := i; j < om.Count-1; j++ {
                om.Employees[j] = om.Employees[j+1]
            }
            om.Count--
            fmt.Printf("Сотрудник %s удален\n", name)
            return
        }
    }
    fmt.Printf("Сотрудник %s не найден\n", name)
}

func (om *OfficeManager) PrintEmployees() {
    fmt.Println("Список сотрудников:")
    for i := 0; i < om.Count; i++ {
        fmt.Printf("%d. %s (%d)\n", i+1, om.Employees[i].Name, om.Employees[i].Age)
    }
    fmt.Println()
}