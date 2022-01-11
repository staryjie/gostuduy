package main

import "fmt"

func defer_demo() {
	defer fmt.Println("Hello")

	fmt.Println("Hi")
}

func mutil_defer() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	fmt.Println("d")
}

func main() {
	// defer_demo()
	mutil_defer()
}
