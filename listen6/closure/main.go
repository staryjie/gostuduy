package main

import (
	"fmt"
)

func sub() func() {
	i := 10
	fmt.Printf("i addr %p\n", &i) // 查看自由变量i的内存地址
	b := func() {
		fmt.Printf("i addr %p\n", &i) // 查看自由变量i的内存地址
		i--
		fmt.Printf("i = %d\n", i) // 查看自由变量i的内存地址
	}
	return b // 返回一个闭包函数
}

func add(base int) func(int) int { // 返回一个闭包函数 func(int) int {}
	return func(i int) int {
		fmt.Printf("base addr %p\n", &base)
		base += 1
		fmt.Printf("base = %d\n", base)

		return base
	}
}

func main() {
	f := sub()
	f()
	f() // 闭包函数复制的是自由变量的指针

	f1 := add(10) // base = 10
	fmt.Println(f1(1))
	fmt.Println(f1(2))

	f2 := add(10)
	fmt.Println(f2(1))
	fmt.Println(f2(2))
}
