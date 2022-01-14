package main

import (
	"fmt"
)

func testPoint1() {
	var a int // 值类型，默认为对应数据类型的默认值
	a = 1000
	fmt.Printf("the address of a: %p, a = %d\n", &a, a)

	var b *int // 指针类型，默认为nil
	fmt.Printf("the address of b: %p, *b = %v\n", &b, b)
	if b == nil {
		fmt.Println("b is nil addr")
	}

	b = &a
	fmt.Printf("the address of b: %p, *b = %v, b = %v\n", &b, *b, b) // *b 解析b指针对应的数据的值
}

func testPoint2() {
	var a int = 200
	var b *int = &a

	fmt.Printf("a = %d\t*b = %d\n", a, *b)

	fmt.Printf("b指针地址存储的值为:%d\n", *b)

	*b = 1000
	fmt.Printf("a = %d\t*b = %d\n", a, *b)

	fmt.Printf("b指针地址存储的值为:%d\n", *b)
}

func modify(p *int) {
	*p = 1000
}

func testPoint3() {
	var a int = 10
	p := &a // p为指针类型
	modify(p)
	fmt.Println(a)
}

func modify_arr(arr *[3]int) {
	(*arr)[0] = 1000
}

func testPoint4() {
	var arr [3]int = [3]int{1, 2, 3}
	modify_arr(&arr) // 传递指针
	fmt.Println(arr)
}

func main() {
	// testPoint1()
	// testPoint2()
	// testPoint3()
	testPoint4()
}
