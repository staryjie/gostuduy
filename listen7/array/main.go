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

func testArray7() {
	// 二维数组
	var arr7 [3][2]int
	fmt.Println(arr7)

	arr7[0][0] = 10
	arr7[0][1] = 11
	arr7[1][0] = 20
	arr7[1][1] = 21
	arr7[2][0] = 30
	arr7[2][1] = 31

	fmt.Println(arr7)

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("arr7[%d][%d] = %d\n", i, j, arr7[i][j])
		}
	}
	fmt.Println()
	for i, e := range arr7 {
		// fmt.Printf("arr7 row[%d] = %v\n", i, e)
		for j, e2 := range e {
			fmt.Printf("arr7[%d][%d] = %d\n", i, j, e2)
		}
	}
}

func testArray8() {
	var a int = 1000
	b := a // 基本数据类型的赋值是值拷贝
	b = 3000
	fmt.Printf("a = %d  b = %d\n", a, b)
}

func modify(arr [3]int) {
	arr[0] = 1000
}

func testArray9() {
	arr9 := [3]int{1, 2, 3}
	brr := arr9 // 数组的拷贝是值拷贝
	brr[0] = 999

	fmt.Printf("arr9 = %v\n", arr9)
	fmt.Printf("brr = %v\n", brr)

	modify(arr9) // 数组传参也是传递值拷贝

	fmt.Printf("arr9 = %v\n", arr9)

}

func main() {
	// testArray1()
	// testArray2()
	// testArray3()
	// testArray4()
	// testArray5()
	// testArray6()
	// testArray7()
	// testArray8()
	testArray9()
}
