package main

import "fmt"

func for_expr() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}
}

func main() {
	for_expr()
}
