package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	a := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(a)
}
