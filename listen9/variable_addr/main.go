package main

import (
	"fmt"
)

func main() {
	var a int
	a = 1000

	fmt.Printf("Address of a: %p, a = %d\n", &a, a)
}
