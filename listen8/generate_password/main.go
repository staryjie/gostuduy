package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "~!@#$%^&*()_+[];':,/?><`|"
)

func init() {
	flag.IntVar(&length, "l", 16, "-l [length] 生成密码的长度")
	flag.StringVar(&charset, "t", "num",
		`-t [charset]
num:     生成只含有数字的密码
char:    生成只含有字母的密码
mix:     生成数字和字母混合的密码
advance: 生成数字、字母、特殊字符混合的密码
	`)
}

func generetePassword() string {
	var password = make([]byte, length, length)
	var sourceStr string

	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}
	// fmt.Printf("sourceStr: %s\n", sourceStr)

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		password[i] = sourceStr[index]
	}

	return string(password)
}

func main() {
	flag.Parse() // 解析参数
	rand.Seed(time.Now().UnixNano())
	// fmt.Printf("length: %d\tcharset: %s\n", length, charset)
	password := generetePassword()
	fmt.Println(password)
}
