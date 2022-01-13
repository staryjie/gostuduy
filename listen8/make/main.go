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

func slicLenCap() {
	a := make([]int, 1, 10)
	a[0] = 10
	fmt.Printf("a = %v\nlen a = %d\ncap a = %d\naddr a = %p\n", a, len(a), cap(a), &a)

	// 切片扩容
	a = append(a, 20)
	fmt.Printf("扩容后:\na = %v\nlen a = %d\ncap a = %d\naddr a = %p\n", a, len(a), cap(a), &a)

	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("扩容后:\na = %v\nlen a = %d\ncap a = %d\naddr a = %p\n", a, len(a), cap(a), &a)
	}

	// 扩容到len = cap之后再次扩容
	a = append(a, 100)
	fmt.Printf("len=cap之后扩容:\na = %v\nlen a = %d\ncap a = %d\naddr a = %p\n", a, len(a), cap(a), &a)
}

func main() {
	// testMake()
	slicLenCap()
}
