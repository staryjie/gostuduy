package main

import (
	"fmt"
)

// 全局变量
var a int = 100

func globalVariable() {
	fmt.Printf("a = %d\n", a)
}

func localVariable() {
	var a int = 1000
	var b int = 99                      // 局部变量,在函数内部
	fmt.Printf("a = %d b = %d\n", a, b) // 局部变量与全局变量同名，局部变量优先级高

	if b == 99 {
		var c int = 98 // 局部变量,在语句块内
		fmt.Printf("c = %d\n", c)
	}

	if d := 100; d > 0 { // 局部变量，在代码块中
		fmt.Printf("d = %d\n", d)
	}

	for i := 0; i < 10; i++ { // 局部变量,在for循环中有效
		fmt.Println(i)
	}
}

func main() {
	// globalVariable()
	localVariable()
}
