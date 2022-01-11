package main

import "fmt"

func defer_demo() {
	defer fmt.Println("Hello")

	fmt.Println("Hi")
}

func main() {
	defer_demo()
}
