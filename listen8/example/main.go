package main

import (
	"fmt"
	"sort"
)

func example1() {
	var sa = make([]string, 5, 10)

	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}

	fmt.Println(sa)
}

func sortArray() {
	a := [5]int{5, 7, 3, 9, 2}
	sort.Ints(a[:])
	fmt.Println(a)

	s := [5]string{"g", "h", "f", "a", "d"}
	sort.Strings(s[:])
	fmt.Println(s)

	f := [5]float64{5.3, 4.7, 9.16, 2.85, 0.13}
	sort.Float64s(f[:])

	fmt.Println(f)
}

func main() {
	// example1()
	sortArray()
}
