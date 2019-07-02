package main

import "fmt"

var a = 5
var z int

func main() {
	fmt.Println(a)
	fmt.Println("z:", z)
	foo()
}

func foo() {
	fmt.Println("hello", a)
	fmt.Printf("%T\n", a)
}
