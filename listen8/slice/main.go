package main

import (
	"fmt"
)

func testSlice1() {
	var a []int // 声明slice
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println(a)
	}

	a[0] = 100
}

func main() {
	testSlice1()
}
