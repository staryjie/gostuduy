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

func main() {
	testArraySum()
}
