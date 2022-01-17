package main

import "fmt"

// 从字符串获取输入 fmt

func testSscanf() {
	var (
		a int
		s string
		f float64
	)
	//var str string = "88 hi 6.8"
	var str string = "88 hi\n\n 6.8"

	//fmt.Sscanf(str,"%d %s %f", &a, &s, &f)
	fmt.Sscanf(str,"%d %s %f\n", &a, &s, &f)
	fmt.Printf("a = %d  s = %s  f = %.2f\n", a, s, f)
}

func testSscan() {
	var (
		a int
		s string
		f float64
	)
	var str string = "88 hi\n\n 6.8"
	fmt.Sscan(str, &a, &s, &f)
	fmt.Printf("a = %d  s = %s  f = %.2f\n", a, s, f)
}

func testSscanln() {
	var (
		a int
		s string
		f float64
	)
	var str string = "88 hi\n\n 6.8"
	fmt.Sscanln(str, &a, &s, &f)
	fmt.Printf("a = %d  s = %s  f = %.2f\n", a, s, f)
}

func main() {
	//testSscanf()
	//testSscan()
	testSscanln()
}