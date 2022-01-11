package main

import (
	"fmt"
)

func justify(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func prime_number() {
	// 打印11-00以内的所有质数
	for i := 2; i < 100; i++ {
		if justify(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	prime_number()
}
