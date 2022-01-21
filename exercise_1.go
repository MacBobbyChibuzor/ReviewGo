package main

import "fmt"

/* Write a function which accepts 2 integers and returns their sum, product and difference. Make a  version with unnamed return variables, and a 2nd version with named return variables.
*/

var (
    numA = 6
    numB = 2
)

func main() {
    
    first := productAndDifference_N(numA, numB)
    Printvalue()
    last := productAndDifference_U(numA, numB)
    Printvalue()
}

func Printvalue() {
    fmt.Println()
}

func productAndDifference_N(x, y int) (sum, product, difference int) {
    sum = x + y
    product = x * y
    difference = x - y
    return
}

func productAndDifference_U(x, y int) (int, int, int) {
    return x + y, x * y, x - y
}