package main

import (
    "fmt"
    "rand"
)

func main() {
    n := rand.Intn(10)
    if n <= 5 {
        fmt.Println("Small")
    } else if n > 5 && n < 10 {
        fmt.Println("Medium")
    } else {
        fmt.Println("Random or large")
    }
    
    if m := rand.Intn(100); m < 10 {
        fmt.Println("Smaller")
    } else if m > 10 && n < 80 {
        fmt.Println("Medium")
    } else {
        fmt.Println("Random or large")
    }
}