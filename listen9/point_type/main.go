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

func testPoint5() {
	var a *int = new(int) // new()给变量分配内存,并返回一个指针
	fmt.Printf("a指针地址存储的值为:%d\n", *a)

	*a = 100
	fmt.Printf("a指针地址存储的值为:%d\n", *a)

	var b *[]int = new([]int)
	fmt.Printf("b指针地址存储的值为:%v\n", *b)
	if *b == nil {
		fmt.Println("*b is nil")
	}
	// nil的时候无法直接操作
	(*b) = make([]int, 3, 5)          // make 为内建类型slice、map和channel分配内存，并可以指定长度和容量 make还能完成初始化
	fmt.Printf("b指针地址存储的值为:%v\n", *b) // [0 0 0]
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("b指针地址存储的值为:%v\n", *b)
}

func modifyInt(a *int) {
	*a = 1000
}

func testPoint6() {
	var a int = 10
	b := a
	modifyInt(&a) // 传递指针

	fmt.Printf("a = %d\tb=%d\n", a, b)
}

func testPoint7() {
	var a int = 10
	var b *int = &a
	var c *int = b
	fmt.Printf("a = %d\t*b=%d\t*c=%d\n", a, *b, *c)

	*c = 100
	fmt.Printf("a = %d\t*b=%d\t*c=%d\n", a, *b, *c)
}

func main() {
	// testPoint1()
	// testPoint2()
	// testPoint3()
	// testPoint4()
	// testPoint5()
	// testPoint6()
	testPoint7()
}
