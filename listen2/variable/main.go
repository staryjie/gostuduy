package main

import "fmt"

func main() {
	var i int
	var b bool
	var s string
	var f float64

	fmt.Printf("i = %d, b = %t, s = %s, f = %.2f\n", i, b, s, f)

	i = 8
	b = true
	s = "hello"
	f = 3.14

	fmt.Printf("i = %d, b = %t, s = %s, f = %.2f\n", i, b, s, f)

	var (
		i1 int
		b1 bool
		s1 string
		f1 float64
	)
	fmt.Printf("i = %d, b = %t, s = %s, f = %.2f\n", i1, b1, s1, f1)
}
