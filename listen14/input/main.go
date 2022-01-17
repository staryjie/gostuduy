package main

import "fmt"

func testScanf() {
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

func testScan() {
	var (
		a int
		s string
		f float64
	)

	fmt.Scan(&a, &s, &f)
	fmt.Printf("a = %d  s = %s  f = %.2f\n", a, s, f)
}

func main() {
	// testScanf()
	testScan()
}
