package main

import "fmt"

func main() {
	var m map[string]int

	m = make(map[string]int, 10)
	m["stu01"] = 1000
	m["stu02"] = 2000
	m["stu03"] = 3000
	m["stu04"] = 4000

	fmt.Printf("m = %v\n", m)

	key := "stu06"
	if result, ok := m[key]; ok {
		fmt.Printf("m[stu03] = %d\n", result)
	} else {
		fmt.Printf("Key %s not esists!", key)
	}
}
