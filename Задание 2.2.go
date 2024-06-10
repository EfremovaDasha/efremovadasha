//Конвертер из римских цифр в арабские

package main

import "fmt"

func romanToInt(s string) int {
    romanValues := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    result := 0
    prev := 0

    for i := len(s) - 1; i >= 0; i-- {
        curr := romanValues[s[i]]
        if curr < prev {
            result -= curr
        } else {
            result += curr
        }
        prev = curr
    }

    return result
}

func main() {
    fmt.Println(romanToInt("MCMXCIV"))
}