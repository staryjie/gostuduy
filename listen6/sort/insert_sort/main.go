package main

import "fmt"

// 插入排序

func insert_sort(arr [9]int) [9]int {
	for i := 1; i < len(arr); i++ { // 从第二个元素开始
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				//fmt.Printf("%d %d --> %d %d  ==> %v\n", arr[j], arr[j-1], arr[j-1], arr[j], arr)
			}
		}
	}

	return arr
}

func main() {
	arr := [9]int{5, 3, 4, 7, 2, 8, 6, 9, 1}
	fmt.Println(arr)
	brr := insert_sort(arr)
	fmt.Println(brr)
}
