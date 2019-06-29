package main

import "fmt"

func main() {
	fmt.Println("hello world")

	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("hello foo")

}

func bar() {
	fmt.Println("And the exited")
}
