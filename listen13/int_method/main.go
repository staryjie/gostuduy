package main

import "fmt"

type Integer int // 自定义类型

func (i Integer) Print() {
	fmt.Println(i)
}

func main() {
	var a Integer = 1000
	a.Print()
}
