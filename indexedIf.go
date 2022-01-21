package main

import "fmt"

func main() {
    str := "Go is a beautiful language!"
    fmt.Println("The length of str is %d.\n", len(str))
    for i := 0; i < len(str); i++ {
        fmt.Println("The character in index %d is %c.", i, str[i])
    }
}