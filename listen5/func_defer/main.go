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

func defer_in_for() {
	for i := 1; i <= 5; i++ {
		defer fmt.Printf("i = %d\n", i)
	}

	fmt.Println("Running...")
	fmt.Println("return")

}

func main() {
	// defer_demo()
	// mutil_defer()
	defer_in_for()
}
