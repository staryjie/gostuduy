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

func main() {
	num := 11
	res := if_expr(num)
	fmt.Println(res)
}
