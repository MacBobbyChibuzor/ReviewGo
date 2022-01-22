package main

import "fmt"

type Address struct {
    Nigeria []string
}

type VCard struct {
    name string
    Address
    birthdate string
}

func myVCard() *VCard{
    Me := VCard{
            name: "Mac"
            Nigeria: []string{"Lagos", "Oyo", "Delta", "Abuja", "Niger", "Kogi", "Osun", "Ogun", "Enugu", "Ekiti", "Edo"}
            birthdate: "21st July"
    }
    fmt.Printf("My VCard is at: %p", &Me)
    return &Me
}

func main() {
    fmt.Printf("%p\n", myVCard
    fmt.Printf("My VCard contains: %p", *myVCard)
}