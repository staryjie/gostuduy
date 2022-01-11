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

func mutil_for() {
	for no, i := 10, 1; i < 10 && no < 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}

func main() {
	// for_expr()
	// break_expr()
	// continue_expr()
	// simpale_for()
	mutil_for()
}
