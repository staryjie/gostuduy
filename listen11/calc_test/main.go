package main

import (
	"fmt"

	"github.com/staryjie/gostudy/listen11/calc"
)

var (
	a, b int = 100, 300
)

func init() {
	fmt.Printf("a = %d\nb = %d\n", a, b)
	fmt.Println("init func in main package")
}

func main() {
	var a, b int = 12, 8

	sum := calc.Add(a, b)
	fmt.Printf("sum = %d\n", sum)
}

/*
	Go语言中代码执行顺序
	按照导入包的先后顺序，逆序执行
	1. 先执行全局变量的初始化
	2. 执行对应的init函数
	3. main包的全局变量初始化
	4. main包init函数
	5. main函数
*/
