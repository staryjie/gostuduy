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

func testArray2() {
	// var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	// arr2 := [5]int{1, 2, 3, 4, 5} // 类型推导初始化
	// arr2 := [...]int{1, 2, 3, 4, 5} // 自动推导数组长度
	arr2 := [5]int{1, 2, 3} // 初始化的时候给一部分元素赋值，其他元素为默认值
	fmt.Println(arr2)
}

func main() {
	// testArray1()
	testArray2()
}
