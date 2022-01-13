package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumArray(arr [10]int) (sum int) {
	// 求数组所有元素的和
	for _, e := range arr {
		sum += e
	}
	return sum
}

func testArraySum() {
	// 随机生成数组
	rand.Seed(time.Now().UnixMilli())
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	fmt.Printf("随机生成的数组为: %v\n", arr)
	sum := sumArray(arr)
	fmt.Printf("sum = %d\n", sum)
}

func targetIndex(arr [5]int, target int) {
	for i := 0; i < len(arr); i++ {
		three := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == three {
				fmt.Printf("arr[%d] + arr[%d] = %d\n", i, j, target)
			}
		}
	}
}

func testSumIndex() {
	arr := [...]int{1, 3, 5, 8, 7}
	targetIndex(arr, 8)
}

func main() {
	// testArraySum()
	testSumIndex()
}

// 数组传参存在的最大的问题：
// 在数组传参的时候传入的数组的长度必须要和接收数组的参数的长度一致，这样对于传参来说是非常不方便的，我们很多时候并不能控制要传递的数组的长度
// 所以数组在生产中并不常用，而是使用切片的频率比较高
