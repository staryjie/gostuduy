package main

import "fmt"

type Test struct { // 结构体的定义
	A int32
	B int32
	C int32
	D int32
}

func main() {
	var t Test
	// struct占用一段连续的内存空间
	fmt.Printf("A addr: %p\n", &t.A)
	fmt.Printf("B addr: %p\n", &t.B)
	fmt.Printf("C addr: %p\n", &t.C)
	fmt.Printf("D addr: %p\n", &t.D)
}
