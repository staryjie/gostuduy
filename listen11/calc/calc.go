package calc

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func init() {
	fmt.Println("init func in package calc")
}
