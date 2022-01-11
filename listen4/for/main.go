package main

import "fmt"

func for_expr() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}
}

func break_expr() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("i=%d\n", i)
	}
}

func continue_expr() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("i=%d\n", i)
	}
}

func simpale_for() {
	var i = 1
	for i <= 10 {
		fmt.Printf("i=%d\n", i)
		i += 2
	}
}

func main() {
	// for_expr()
	// break_expr()
	// continue_expr()
	simpale_for()
}
