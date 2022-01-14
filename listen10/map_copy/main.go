package main

import (
	"fmt"
)

func modify(m map[string]int) {
	m["stu999"] = 999
}

func main() {
	var m map[string]int

	m = make(map[string]int, 10)
	m["stu01"] = 1000
	m["stu02"] = 2000
	m["stu03"] = 3000
	m["stu04"] = 4000

	fmt.Printf("begin m = %v\n", m)

	m2 := m
	m["stu03"] = 333
	fmt.Printf("after modify m2['stu03'] m = %v\n", m)
	fmt.Printf("m2 = %v\n", m2)

	modify(m)

	fmt.Printf("After modify m m2 = %v\n", m2)

}
