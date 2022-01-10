package main

import "fmt"

func main() {
	const i int = 100
	const s = "hello"

	fmt.Printf("i = %d, s = %s\n", i, s)

	const (
		i1 int = 101
		s1     = "hello"
	)

	fmt.Printf("i1 = %d, s1 = %s\n", i1, s1)

	const (
		a = 100
		b
		c = 200
		d
	)

	fmt.Println(a, b, c, d)

	const (
		one = iota + 1
		two
		three
		four
	)

	fmt.Println(one, two, three, four)

	const (
		_  = 1 << (10 * iota)
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
	)

	fmt.Println(KB, MB, GB, TB)
}
