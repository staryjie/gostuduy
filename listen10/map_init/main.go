package main

import (
	"fmt"
)

func main() {
	var m map[string]int = map[string]int{ // map初始化，map必须初始化才能使用
		"stu01": 1000,
		"stu02": 2000,
		"stu03": 3000,
	}

	fmt.Println(m)

	// 修改map中的元素
	m["stu01"] = 1111

	// 插入数据
	m["stu04"] = 4000

	// 访问map中的元素
	fmt.Printf("the value of m[stu04] is: %d\n", m["stu04"])
}
