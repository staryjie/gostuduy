package main

import "fmt"

func mutil_args(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b

	return
}

func calc(n ...int) (sum int) {
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return
}

func main() {
	// sum, sub := mutil_args(12, 6)
	// fmt.Printf("sum = %d  sub = %d\n", sum, sub)

	// sum2, _ := mutil_args(19, 3)
	// fmt.Printf("sum = %d\n", sum2)

	sum := calc(1, 2, 3, 4, 5)
	fmt.Printf("sum = %d\n", sum)
}
