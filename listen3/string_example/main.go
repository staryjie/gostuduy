package main

import "fmt"

// 字符串逆序 - 中文会乱码
func ReverseString(s string) string {
	bs := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		tmp := bs[len(s)-i-1]
		bs[len(s)-i-1] = bs[i]
		bs[i] = tmp
	}
	return string(bs)
}

func main() {
	s := "hello"
	res := ReverseString(s)
	fmt.Println(res)
}
