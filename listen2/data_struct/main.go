package main

// 数据类型

import "fmt"

func main() {
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
