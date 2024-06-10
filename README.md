Практика

Лекция 1.

Задание №1.
Написать одну из трёх сортировок статического массива (сортировку вставками, пузырьком, слиянием)

Реализация сортировки слиянием:

package main

import "fmt"

func main() {
    nums := []int{5, 2, 9, 1, 7, 3, 8, 4, 6}
    mergeSort(nums)
    fmt.Println(nums)
}

func mergeSort(arr []int) {
    if len(arr) <= 1 {
        return
    }

    mid := len(arr) / 2
    left, right := arr[:mid], arr[mid:]
    mergeSort(left)
    mergeSort(right)

    merge(arr, left, right)
}

func merge(arr, left, right []int) {
    i, j, k := 0, 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }

    for i < len(left) {
        arr[k] = left[i]
        i++
        k++
    }

    for j < len(right) {
        arr[k] = right[j]
        j++
        k++
    }
}

Вывод: [1 1 1 1 3 3 4 4 6]

Задание №2.
Написать программу которая хранит данные об офисных сотрудниках, сделать операцию добавления и удаления, максимум сколько можно добавить сотрудников 512.

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

Вывод:
Список сотрудников:
1. John Doe (30)
2. Jane Smith (35)
3. Bob Johnson (40)

Сотрудник Jane Smith удален
Список сотрудников:
1. John Doe (30)
2. Bob Johnson (40)


