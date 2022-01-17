package main

import "fmt"

func testInput() {
	var (
		a int
		s string
		f float64
	)

	// fmt.Scanf("%d %s %f", &a, &s, &f)
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%f\n", &f)

	fmt.Printf("a = %d  s = %s  f = %.2f\n", a, s, f)
}

func main() {
	testInput()
}
