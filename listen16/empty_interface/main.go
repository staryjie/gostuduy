package main

import "fmt"

func describe(a interface{}) {
	fmt.Printf("%T %v\n", a, a)
}

type Student struct {
	Name string
	Sex int
}

func main() {
	var a interface{}  // 空接口,空接口可以存放任何数据类型

	var b int = 64
	a = b

	describe(a)

	var s string = "hello"
	a = s
	describe(a)

	var m map[string]int = make(map[string]int, 100)
	m["hello"] = 100
	a = m
	describe(a)

	stu := Student{Name: "Tom", Sex: 1}
	describe(stu)
}
