package main

import (
	"fmt"
)

func add(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	a := 8
	b := 6
	sum, sub := add(a, b)

	fmt.Print(sum, sub)
}
