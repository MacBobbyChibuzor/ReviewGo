package main

import "fmt"

/*
var a int = 10
pointertoA = &a 
*/

func main() {
    /*
    fmt.Println(a)
    fmt.Println(pointertoA)
    */
    x := 5
    var pointertoX *int
    pointertoX = &x
    fmt.Println(pointertoX)
}