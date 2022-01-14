package main

import (
	"fmt"
)

func main() {
	var m map[string]int = make(map[string]int, 10)
	m["stu01"] = 1000
	m["stu02"] = 2000
	m["stu03"] = 3000
	m["stu04"] = 4000

	fmt.Printf("m = %v\n", m)

	key := "stu01"
	delete(m, key)
	fmt.Printf("m = %v\n", m)
}
