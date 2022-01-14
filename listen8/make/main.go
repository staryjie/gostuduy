package main

import (
	"fmt"
)

func testMake() {
	a := make([]int, 5, 10) // []type len cap

	a[0] = 10
	a[1] = 20
	fmt.Printf("a = %v\n", a)
}

func slicLenCap() {
	a := make([]int, 1, 10)
	a[0] = 10
	fmt.Printf("a = %v\nlen a = %d\ncap a = %d\naddr a = %p\n", a, len(a), cap(a), &a)

	// 切片扩容
	a = append(a, 20)
	fmt.Printf("扩容后:\na = %v\nlen a = %d\ncap a = %d\naddr a = %p\n", a, len(a), cap(a), &a)

	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("扩容后:\na = %v\nlen a = %d\ncap a = %d\naddr a = %p\n", a, len(a), cap(a), &a)
	}

	// 扩容到len = cap之后再次扩容
	a = append(a, 100)
	fmt.Printf("len=cap之后扩容:\na = %v\nlen a = %d\ncap a = %d\naddr a = %p\n", a, len(a), cap(a), &a)
}

func testCap() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	b := a[1:3]
	fmt.Printf("len a = %d\ncap b = %d\n", len(b), cap(b)) // len = 2 cap = 7
	// a的长度为8，a[1:3]只有两个元素，a[0]并不在b能操作的范围，所以cap = 8 -1 = 7

	b = b[:cap(b)]
	fmt.Printf("再切片：len b = %d\ncap b = %d\n", len(b), cap(b))
}

func testSlice() {
	var a []int
	fmt.Printf("addr a = %p  len a = %d  cap d = %d\n", &a, len(a), cap(a))

	if a == nil {
		fmt.Println("a is nil")
	}

	// a[0] = 100  // index out of range [0] with length 0  空切片无法直接访问

	// 空切片可以通过append追加元素
	a = append(a, 100)
	fmt.Printf("addr a = %p  len a = %d  cap d = %d\n", &a, len(a), cap(a)) // 1 1

	a = append(a, 200)
	fmt.Printf("addr a = %p  len a = %d  cap d = %d\n", &a, len(a), cap(a)) // 2 2

	a = append(a, 300)
	fmt.Printf("addr a = %p  len a = %d  cap d = %d\n", &a, len(a), cap(a)) // 3 4

	a = append(a, 400)
	fmt.Printf("addr a = %p  len a = %d  cap d = %d\n", &a, len(a), cap(a)) // 4 4

	a = append(a, 500)
	fmt.Printf("addr a = %p  len a = %d  cap d = %d\n", &a, len(a), cap(a)) // 5 8
}

func testSubSliceCap() {
	a := make([]int, 5, 5)
	fmt.Printf("addr a = %p  len a = %d  cap d = %d\n", &a, len(a), cap(a))

	b := a[0:3]
	fmt.Printf("a = %v\t b = %v\n", a, b)
	fmt.Printf("addr b = %p  len b = %d  cap b = %d\n", &b, len(b), cap(b))

	b = append(b, 100)
	fmt.Printf("a = %v\t b = %v\n", a, b)
	fmt.Printf("addr b = %p  len b = %d  cap b = %d\n", &b, len(b), cap(b))

	b = append(b, 200)
	fmt.Printf("a = %v\t b = %v\n", a, b)
	fmt.Printf("addr b = %p  len b = %d  cap b = %d\n", &b, len(b), cap(b))

	b = append(b, 300)
	fmt.Printf("a = %v\t b = %v\n", a, b)
	fmt.Printf("addr b = %p  len b = %d  cap b = %d\n", &b, len(b), cap(b))

	b = append(b, 400)
	fmt.Printf("a = %v\t b = %v\n", a, b)
	fmt.Printf("addr b = %p  len b = %d  cap b = %d\n", &b, len(b), cap(b))
}

func testAppend() {
	a := []int{1, 3, 5}
	b := []int{7, 9, 11}

	a = append(a, b...) // 将一个切片追加到另一个切片中
	fmt.Printf("a = %v\n", a)
}

func sumArray(arr []int) (sum int) {
	for _, e := range arr {
		sum += e
	}

	return
}

func testSumArray() {
	a := [3]int{1, 2, 3}
	suma := sumArray(a[:])
	fmt.Println(suma)

	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumb := sumArray(b[:])
	fmt.Println(sumb)
}

func modifySlice(arr []int) {
	arr[0] = 1000
}

func testModifySlice() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	modifySlice(a[:])
	fmt.Println(a)
}

func copySlice() {
	// a := []int{1, 2, 3}
	a := []int{1}
	b := []int{4, 5, 6}

	copy(a, b) // 将b拷贝到a切片，拷贝元素个数取决于目标切片的长度，不会自动扩容

	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	// testMake()
	// slicLenCap()
	// testCap()
	// testSlice()
	// testSubSliceCap()
	// testAppend()
	// testSumArray()
	// testModifySlice()
	copySlice()
}
