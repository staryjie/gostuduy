package main

import (
	"fmt"
)

func if_expr(num int) string {
	if num%2 == 0 {
		return fmt.Sprintf("num %d is even number.", num)
	} else {
		return fmt.Sprintf("num %d is odd number.", num)
	}
}

func if_else_if_expr(num int) {
	if num > 5 && num <= 10 {
		fmt.Println("5 < num <= 10")
	} else if num > 0 && num <= 5 {
		fmt.Println("0 < num <= 5")
	} else {
		fmt.Printf("num is %d\n", num)
	}
}

func main() {
	num := 5

	// if else 语句
	/*
		res := if_expr(num)
		fmt.Println(res)
	*/

	// if else if else
	if_else_if_expr(num)
}
