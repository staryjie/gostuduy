package main

import (
	"fmt"
)

// 全局变量
var a int = 100

func globalVariable() {
	fmt.Printf("a = %d\n", a)
}

func main() {
	globalVariable()
}
