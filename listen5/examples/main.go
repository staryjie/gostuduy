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

func is_daffodils(n int) bool {
	first := n % 10
	second := (n / 10) % 10
	third := (n / 100) % 10

	if first*first*first+second*second*second+third*third*third == n {
		return true
	}
	return false
}

func daffodils_number() {
	// 100-999之间是水仙花数
	for i := 100; i < 1000; i++ {
		if is_daffodils(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	// prime_number()
	daffodils_number()
}
