package main

import "fmt"

func main() {
    var (
        min int
        max int
        )
    min, max = MinMax(70, 85)
    fmt.Printf("The mininum is %d, and the maximum is %d.", min, max)
}
func MinMax(a, b int) (min, max int) {
    if a < b {
        min = a
        max = b
    } else {
        min = b
        max = a
    }
}