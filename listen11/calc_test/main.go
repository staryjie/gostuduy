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
