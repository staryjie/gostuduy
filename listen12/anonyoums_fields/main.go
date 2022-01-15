package main

import (
	"fmt"
)

type User struct {
	Username string
	Age      int
	Sex      int
	int      // 匿名字段直接使用数据类型作为变量
	string
}

func main() {
	var user User
	user.Username = "Tom"
	user.Age = 18
	user.Sex = 1
	user.int = 100
	user.string = "Hi"

	fmt.Printf("user = %#v\n", user)
}
