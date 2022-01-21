package main

import "fmt"

func main() {
    f := division(10, 2)
    fmt.Printf("Division is %d.", f)
    m := multiply(2, 3, 4)
    fmt.Printf("Multiplication is %d.", m)
}

func division(num, den int) int {
    if den == 0 {
        return 0
    }
    return num / den
}

func multiply(a, b, c int) int {
    return a * b * c
}