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