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

func testSlice2() {
	a := [5]int{76, 77, 78, 79, 80} // 数组
	var b []int
	b = a[1:4]

	fmt.Println(b)
}

func main() {
	// testSlice1()
	testSlice2()
}
