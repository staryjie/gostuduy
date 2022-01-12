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
	var b int = 99 // 局部变量
	fmt.Printf("b = %d\n", b)
}

func main() {
	// globalVariable()
	localVariable()
}
