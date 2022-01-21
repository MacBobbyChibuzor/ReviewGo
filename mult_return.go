package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
    numx2, numx3 = getx2andx3(num)
    Printout()
    numx2, numx3 = getx2andx3_2(num)
    Printout()
}

func Printout() {
    fmt.Println("num = %d, 2 x num = %d, 3 x num = %d", num, numx2, numx3)
}

func getx2andx3(input int) (int, int) {
    return 2 * input, 3 * input
}

func getx2andx3_2(input int) (x2, x3 int) {
    x2 = 2 * input
    x3 = 3 * input
    return x2, x3
}