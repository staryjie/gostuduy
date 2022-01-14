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

func main() {
	testPoint1()
}
