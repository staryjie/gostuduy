package main

import "fmt"

// 从终端获取用户输入  fmt

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

func testScanln() {
	var (
		a int
		s string
		f float64
	)
	fmt.Scanln(&a, &s, &f)
	fmt.Printf("a = %d  s = %s  f = %.2f\n", a, s, f)
}

func main() {
	// testScanf()
	// testScan()
	testScanln()
}
