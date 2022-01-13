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
	fmt.Printf("b[0] = %d\n", b[0])
	fmt.Printf("b[1] = %d\n", b[1])
	fmt.Printf("b[2] = %d\n", b[2])
}

func testSlice3() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
}

func testSlice4() {
	// 切片的常用操作
	a := [5]int{76, 77, 78, 79, 80} // 数组
	var b []int
	b = a[1:4]

	fmt.Printf("b = %v\n", b)

	c := a[1:]
	fmt.Printf("c = %v\n", c)

	d := a[:3]
	fmt.Printf("d = %v\n", d)

	e := a[:]
	fmt.Printf("e = %v\n", e)
}

func main() {
	// testSlice1()
	// testSlice2()
	// testSlice3()
	testSlice4()
}
