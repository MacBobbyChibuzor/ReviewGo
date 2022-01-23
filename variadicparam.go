package main

import "fmt"

func main() {
	varPara(x)
}

var x = [...]int{2,3,4,5,6,7,8,9}

func varPara(x ...int) {
	fmt.Println(x)
}