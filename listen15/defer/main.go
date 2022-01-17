package main

import "fmt"

func funcA() int {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func funcB() (x int) {
	defer func() {
		x += 1
	}()

	return 5
}

func funcC() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func funcD() (x int) {
	defer func(x int) {
		x += 1
	}(x)
	return 5
}

func main() {
	fmt.Println(funcA())
	fmt.Println(funcB())
	fmt.Println(funcC())
	fmt.Println(funcD())
}

/*
	defer执行原理
	1. defer在函数返回之前被调用
	2. return的实现
		2.1 设置返回值变量的值
		2.2 执行RET指令
	3. defer原理
		3.1 函数返回前设置返回变量的值
		3.2 调用defer语句
		3.3 执行RET指令
*/
