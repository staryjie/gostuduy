package main

import "fmt"

func testString() {
	var s string = "hello" // 字符串底层就是byte数组
	fmt.Printf("s[0] = %c len(s) = %d\n", s[0], len(s))

	// 遍历字符串
	for i, e := range s {
		fmt.Printf("s[%d] = %c\n", i, e)
	}

	// 字符串不能直接修改
	// s[0] = 'i'
	// fmt.Println(s)

	var bs []byte
	bs = []byte(s)
	bs[0] = 'i' // 要使用单引号
	// fmt.Println(bs)
	s = string(bs) // 字符串可以先转成p[]byte然后修改，修改完再转为字符串
	fmt.Println(s)

	// 字符串长度就是底层[]byte的长度
	fmt.Printf("len(s) = %d   len(bs) = %d\n", len(s), len(bs))

	var s1 = "Hi,中国" // 一个中文字符占3个字节
	fmt.Printf("len(s1) = %d\n", len(s1))

	var s2 rune = '中'
	fmt.Printf("s2 = %c\n", s2)

	var s3 = "Hi,中国"
	var rs []rune
	rs = []rune(s3)

	fmt.Printf("len(rs) = %d\n", len(rs))
}

func main() {
	testString()
}
