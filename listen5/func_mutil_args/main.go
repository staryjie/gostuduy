package main

import "fmt"

func mutil_args(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b

	return
}

func main() {
	sum, sub := mutil_args(12, 6)
	fmt.Printf("sum = %d  sub = %d\n", sum, sub)
}
