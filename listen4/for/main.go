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

func main() {
	// for_expr()
	// break_expr()
	continue_expr()
}
