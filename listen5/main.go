package main

import "fmt"

func sayHi() {
	fmt.Println("Hello World!")
}

func add(a, b int) int {
	return a + b
}

func main() {
	// sayHi()
	a, b := 3, 4
	sum := add(a, b)
	fmt.Println(sum)
}
