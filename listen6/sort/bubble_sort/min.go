package main

import "fmt"

// 冒泡排序

func bubble_sort(arr [9]int) [9]int {
	for i := len(arr) - 1; i >= 0; i-- { // 最后一个元素开始
		for j := i - 1; j >= 0; j-- { // 倒数第二个元素开始
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i] // 对比大小，符合条件的交换位置，然后往前推一位
			}
		}
	}
	return arr
}

func main() {
	arr := [9]int{5, 3, 4, 7, 2, 8, 6, 9, 1}
	fmt.Println(arr)
	brr := bubble_sort(arr)
	fmt.Println(brr)
}
