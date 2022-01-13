package main

import (
	"fmt"
)

func testMake() {
	a := make([]int, 5, 10) // []type len cap

	a[0] = 10
	a[1] = 20
	fmt.Printf("a = %v\n", a)
}

func main() {
	testMake()
}
