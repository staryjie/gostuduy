package main

import (
	"fmt"

	"github.com/staryjie/gostudy/listen12/user"
)

func main() {
	var user1 user.User
	user1.Age = 18
	fmt.Printf("user1 = %#v\n", user1)

	// 通过构造函数实现结构体实例初始化
	user2 := user.NewUser("Tom", "http://www.baidu.com", 18, 1)
	fmt.Printf("user2 = %#v\n", user2)
}
