package main

// 数据类型

import (
	"fmt"
	"strings"
)

func bool_demo() {
	// 布尔型
	var b bool // 没有赋值，默认值为false
	fmt.Println(b)

	b = true // 赋值
	fmt.Println(b)

	b = !b // 取反
	fmt.Println(b)

	// && 与操作
	var a bool = true

	if a == true && b == true {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

	// || 或操作
	if a == true || b == true {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

	// 格式化输出
	fmt.Printf("a = %t  b = %t\n", a, b)
}

func int_demo() {
	var a int8
	a = -128
	fmt.Printf("a = %d\n", a)
	a = 127
	fmt.Printf("a = %d\n", a)

	var b int
	b = 123456
	fmt.Printf("b = %d\n", b)
}

func float_demo() {
	var f1 float32
	var f2 float64

	fmt.Println(f1, f2)

	f1 = 5.6
	f2 = 6.18
	fmt.Println(f1, f2)

	i := int(f2) // 类型转换
	fmt.Println(i)

	// 格式化输出
	fmt.Printf("%d %.2f %x\n", i, f2, i)
}

func string_demo() {
	var s string

	fmt.Println(s) // 默认值是空字符串

	s = "hello"
	fmt.Println(s)

	s1 := "world"

	fmt.Println(s1)

	// 格式化输出
	fmt.Printf("s = %s s1 = %v\n", s, s1) // %v 万能输出占位符

	// 反引号
	s3 := `hello
world\t
this
is\n
a
string`

	fmt.Println(s3)
}

func oper_string() {
	s1 := "hello"

	fmt.Printf("the len of s1 is %d\n", len(s1)) // len() 求字符串长度

	// 字符串拼接
	s2 := " world"
	fmt.Println(s1 + s2)

	newS := fmt.Sprint(s1, s2)
	fmt.Println(newS)

	bs := strings.Builder{}
	bs.WriteString(s1)
	bs.WriteString(s2)

	fmt.Println(bs.String())

	// sl := []string{"hello", "world"}
	sl := []string{s1, s2}
	fmt.Println(strings.Join(sl, ""))

	// 字符串分割
	ipaddr := "192.168.5.1;192.168.5.231;192.168.5.74"
	ips := strings.Split(ipaddr, ";")
	fmt.Println(ips)
	fmt.Println(ips[0])
	fmt.Println(ips[1])
	fmt.Println(ips[2])

	// 字符串包含
	fmt.Println(strings.Contains(ipaddr, "192.168.5.1"))

	// 判断字符串开头结尾
	fmt.Println(strings.HasPrefix(ipaddr, "192"))
	fmt.Println(strings.HasSuffix(ipaddr, "74"))

	// 寻找元素索引
	fmt.Println(strings.Index(ipaddr, "231"))
	fmt.Println(strings.Index(ipaddr, "wqwq")) // -1

	fmt.Println(strings.LastIndex(ipaddr, "5"))

}

func main() {
	// bool_demo()
	// int_demo()
	// float_demo()
	// string_demo()
	oper_string()
}
