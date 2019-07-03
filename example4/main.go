package main

import "fmt"

type x int

var a x
var y int

func main() {
	y := int(a)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	//fmt.Println(a)
	fmt.Printf("%T\n", a)
	a = 42
	fmt.Println(a)
}
