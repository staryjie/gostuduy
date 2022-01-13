package main

import (
	"fmt"
)

func testArray1() {
	var arr1 [5]int // 定义数组
	arr1[0] = 200   // 给数组的元素赋值
	arr1[1] = 300
	fmt.Println(arr1)
}

func main() {
	testArray1()
}
