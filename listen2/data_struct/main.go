package main

// 数据类型

import "fmt"

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

func main() {
	// bool_demo()
	// int_demo()
	float_demo()
}
