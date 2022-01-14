package main

import (
	"fmt"
)

func main() {
	// 空map nil map
	var m map[string]int // map为引用类型，声明不初始化的话为nil
	if m == nil {        // nil的map无法直接访问和操作，必须要先初始化
		m = make(map[string]int, 10) // 初始化map，指定长度
	}
	m["stu01"] = 1000
	m["stu02"] = 2000
	m["stu03"] = 3000
	m["stu04"] = 4000

	fmt.Printf("m = %v\n", m)
}
