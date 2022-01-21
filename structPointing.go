package main

import "fmt"
import "strings"

type Person struct {
    firstName string
    lastName  string
}

func upPerson(pointertoPerson *Person) {
    ptrFName := strings.ToUpper("ptrFName")
    ptrLName := strings.ToUpper("ptrLName")
}

func main() {
    // first case with value type
    var human Person
    human.firstName = "Mac"
    human.lastName  = "Uzor"
    upPerson(&human)
    fmt.Println("Name's %s %s", human.firstName, human.lastName)
    
    // second case with pointer
    man := new(Person)
    man.firstName = "Chris"
    man.lastName  = "Evans"
    upPerson(man)
    fmt.Println("Other guy name's %s %s", man.firstName, man.lastName)
    
    // third case with literals
    ptrToPerson := &Person{
        "Henry"
        "Ford"
    }
    upPerson(ptrToPerson)
    fmt.Printf("The third guy's %s %s.", ptrToPerson.firstName, ptrToPerson.lastName)
}