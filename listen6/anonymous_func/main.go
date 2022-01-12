package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func anonymous() {
	f1 := add
	fmt.Printf("%T\n", f1)
	sum := f1(3, 4)
	fmt.Printf("3 + 4 = %d\n", sum)
}

func anonymous_demo() {
	f := func(a, b int) int {
		return a + b
	}
	fmt.Printf("%T\n", f)

	sum := f(2, 5)
	fmt.Printf("2 + 5 = %d\n", sum)
}

func defer_anonymous() {
	i := 100
	defer fmt.Printf("defer: i = %d\n", i)
	defer func() {
		fmt.Printf("defer func: i = %d\n", i)
	}()
	i = 200
	fmt.Println(i)

	return
}

func main() {
	//anonymous()
	// anonymous_demo()
	defer_anonymous()
}
