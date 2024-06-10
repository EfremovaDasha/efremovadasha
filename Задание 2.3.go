//Заполнение двумерного массива случайными уникальными числами

package main

import (
    "fmt"
    "math/rand"
)

func fillUnique2DArray(rows, cols int) [][]int {
    arr := make([][]int, rows)
    for i := range arr {
        arr[i] = make([]int, cols)
    }

    used := make(map[int]bool)
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            num := rand.Intn(rows * cols)
            for used[num] {
                num = rand.Intn(rows * cols)
            }
            used[num] = true
            arr[i][j] = num
        }
    }

    return arr
}

func main() {
    arr := fillUnique2DArray(3, 4)
    for _, row := range arr {
        fmt.Println(row)
    }
}