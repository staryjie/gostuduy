package main

import (
	"fmt"
	"strings"
	"time"
)

/*
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
*/

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func testClosure1() {
	f := Adder()

	ret := f(1) // x = 0; 0 + 1 = 1
	fmt.Printf("f(1): ret = %d\n", ret)

	ret = f(20) // x = 1; 1 + 20 = 21
	fmt.Printf("f(20): ret = %d\n", ret)

	ret = f(100) // x = 21; 21 + 100 = 121
	fmt.Printf("f(100): ret = %d\n", ret)
	// 闭包函数引用的自由变量，在同一个闭包函数实例的多次调用中都生效

	f1 := Adder()
	ret = f1(1) // x = 0; 0 + 1 = 1
	fmt.Printf("f1(1): ret = %d\n", ret)

	ret = f1(1000) // x = 1; 1 + 1000 = 1001
	fmt.Printf("f1(1000): ret = %d\n", ret)
}

func testClosure2() {
	tmp1 := add(10)      // base = 10
	fmt.Println(tmp1(1)) // base = 10; 10 + 1 = 11
	fmt.Println(tmp1(2)) // base = 11; 11 + 2 = 13

	tmp2 := add(100)     // base = 100
	fmt.Println(tmp2(1)) // base = 100; 100 + 1 = 101
	fmt.Println(tmp2(2)) // base = 101; 101 + 2 = 103
}

func makeSuffixFunc(suffix string) func(string) string { // 返回一个 func(string) string {} 匿名函数作为闭包
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func testMakeSuffix() {
	sfx1 := makeSuffixFunc(".bmp")
	sfx2 := makeSuffixFunc(".jpg")

	fmt.Println(sfx1("test.bmp"))
	fmt.Println(sfx2("test"))
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func testCalc() {
	add, sub := calc(10) // base = 10
	fmt.Println(add(1))  // base = 10; 10 + 1 = 11
	fmt.Println(sub(2))  // base = 11; 11 - 2 = 9
	fmt.Println(add(3))  // base = 9; 9 + 3 = 12
	fmt.Println(sub(4))  // base = 12; 12 - 4 = 8
	fmt.Println(add(5))  // base = 8; 8 + 5 = 13
	fmt.Println(sub(6))  // base = 13; 13 - 6 = 7
	fmt.Println(add(7))  // base = 7; 7 + 7 = 14
	fmt.Println(sub(8))  // base = 14; 14 - 8 = 6
}

func go_closure() {
	for i := 1; i < 5; i++ {
		go func(index int) {
			// go func() {
			fmt.Println(index)
			// fmt.Println(i)
			// }()
		}(i)
		time.Sleep(1 * time.Second)
	}
}

func defer_closure() {
	for i := 1; i <= 10; i++ {
		defer func(index int) {
			time.Sleep(time.Second)
			fmt.Println(index)
		}(i)
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	/*
		f := sub()
		f()
		f() // 闭包函数复制的是自由变量的指针

		f1 := add(10) // base = 10
		fmt.Println(f1(1))
		fmt.Println(f1(2))

		f2 := add(10)
		fmt.Println(f2(1))
		fmt.Println(f2(2))
	*/

	// testClosure1()
	// testClosure2()
	// testMakeSuffix()
	// testCalc()
	// go_closure()
	defer_closure()
}
