package main

import(
    "fmt"
    "strings"
)

func main() {
    var sample string = "An example string"
    fmt.Printf("Does the string \"%s" have a prefix \"%s?", sample, "An")
    fmt.Printf("%t\n", strings.HasPrefix(sample, "An"))
}