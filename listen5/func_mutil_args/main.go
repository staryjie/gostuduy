package main

import "fmt"

func mutil_args(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b

	return
}

func calc(a int, n ...int) (sum int) {
	sum += a
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

	// sum := calc(1, 2, 3, 4, 5)
	sum := calc(8, 1, 2, 3, 4, 5)  // a int 是必传的参数，至少需要传递一个参数
	fmt.Printf("sum = %d\n", sum)
}
