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

func ReverseStringCn(s string) string {
	var rs []rune = []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		tmp := rs[len(rs)-i-1]
		rs[len(rs)-i-1] = rs[i]
		rs[i] = tmp
	}
	return string(rs)
}

func Huiwen(s string) bool {
	rs := []rune(s)

	for i := 0; i < len(rs)/2; i++ {
		tmp := rs[len(rs)-i-1]
		rs[len(rs)-i-1] = rs[i]
		rs[i] = tmp
	}
	s2 := string(rs)

	if s == s2 {
		return true
	} else {
		return false
	}

}

func main() {
	s := "上海自来水来自海上"
	// res := ReverseString(s)
	// res := ReverseStringCn(s)
	// fmt.Println(res)
	fmt.Println(Huiwen(s))
}
