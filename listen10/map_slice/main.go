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

func main() {
	sliceMap()
}
