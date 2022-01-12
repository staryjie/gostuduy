package main

import (
	"fmt"
)

// 选择排序

func select_sort(arr [9]int) [9]int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
				//fmt.Println(arr)
			}
		}
	}
	return arr
}

func main() {
	arr := [9]int{5, 3, 4, 7, 2, 8, 6, 9, 1}
	fmt.Println(arr)
	brr := select_sort(arr)
	fmt.Println(brr)
}
