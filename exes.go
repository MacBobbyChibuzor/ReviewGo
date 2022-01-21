package main

import(
    "fmt"
)

func main() {
    emm(a)
    fmt.Println("... is different from")
    enn(b)
}

func emm(m string) string{
    return fmt.Print("Emm func has %s", m)
}

func enn(n string) string{
    return fmt.Print("Enn func has %s", n)
}