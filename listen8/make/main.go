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

func testCap() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	b := a[1:3]
	fmt.Printf("len a = %d\ncap b = %d\n", len(b), cap(b)) // len = 2 cap = 7
	// a的长度为8，a[1:3]只有两个元素，a[0]并不在b能操作的范围，所以cap = 8 -1 = 7

	b = b[:cap(b)]
	fmt.Printf("再切片：len b = %d\ncap b = %d\n", len(b), cap(b))
}

func main() {
	// testMake()
	// slicLenCap()
	testCap()
}
