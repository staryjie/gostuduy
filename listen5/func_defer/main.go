package main

import "fmt"

func defer_demo() {
	defer fmt.Println("Hello")

	fmt.Println("Hi")
}

func mutil_defer() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	fmt.Println("d")
}

func defer_in_for() {
	for i := 1; i <= 5; i++ {
		defer fmt.Printf("i = %d\n", i)
	}

	fmt.Println("Running...")
	fmt.Println("return")

}

func defer_assign() {
	// defer语句中使用变量，编译时就完成赋值了
	var a int = 0
	defer fmt.Printf("defer: a = %d\n", a) // 编译到这里的时候，a=0， 直接赋值

	a = 1000
	fmt.Printf("a = %d\n", a)
}

func main() {
	// defer_demo()
	// mutil_defer()
	// defer_in_for()
	defer_assign()
}
