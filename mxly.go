package main

import (
    "fmt"
)

type Person struct{
    Name string
    Age  int
    ID   int
}

func (p Person) Props() string {
    return fmt.Printf("Hi, I am %s and I am %d years old. My ID is %d.", p.Name, p.Age, p.ID)
}

func main() {
    p := Person{
        Name: "Mac",
        Age:  120,
        ID:   20191750,
    }
    p.Props()
}