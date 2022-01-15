package main

import (
	"fmt"
)

func sliceMap() {
	s := make([]map[string]int, 5, 16)

	for i, e := range s {
		fmt.Printf("s[%d] = %v\n", i, e)
	}

	fmt.Println()

	s[0] = make(map[string]int, 16)
	s[0]["stu01"] = 1000
	s[0]["stu02"] = 2000
	s[0]["stu03"] = 3000
	s[0]["stu04"] = 4000

	for i, e := range s {
		fmt.Printf("s[%d] = %v\n", i, e)
	}

}

func mapSlice() {
	m := make(map[string][]int, 16)

	key := "stu01"

	sliceSource, ok := m[key]
	if !ok {
		m[key] = make([]int, 0, 16)
		sliceSource = m[key]
	}

	sliceSource = append(sliceSource, 100)
	sliceSource = append(sliceSource, 200)
	sliceSource = append(sliceSource, 300)
	sliceSource = append(sliceSource, 400)

	m[key] = sliceSource

	fmt.Printf("m = %v\n", m)
}

func main() {
	// sliceMap()
	mapSlice()
}
