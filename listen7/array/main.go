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
	// arr2 := [5]int{1, 2, 3} // 初始化的时候给一部分元素赋值，其他元素为默认值
	arr2 := [5]int{1: 5, 3: 8} // 给指定索引的元素赋值
	fmt.Println(arr2)
}

func testArray3() {
	var arr3 [5]int = [5]int{1, 2, 3, 4, 5}
	var brr [5]int
	brr = arr3

	brr[3] = 8
	fmt.Printf("arr3 = %v\n", arr3)
	fmt.Printf("brr = %v\n", brr)
}

func testArray4() {
	arr4 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("len arr4 = %d\n", len(arr4))
}

func testArray5() {
	// for循环遍历数组
	arr5 := [5]int{2: 400, 4: 350}
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("arr5[%d] = %d\n", i, arr5[i])
	}
}

func testArray6() {
	// for range遍历数组
	arr6 := [5]int{1, 2, 3, 4, 5}
	for i, e := range arr6 {
		fmt.Printf("arr6[%d] = %d\n", i, e)
	}
}

func main() {
	// testArray1()
	// testArray2()
	// testArray3()
	// testArray4()
	// testArray5()
	testArray6()
}
