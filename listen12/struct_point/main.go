package main

import "fmt"

type User struct { // 结构体的定义
	Username string
	Sex      int
	Age      int
	AvataUrl string
}

func main() {
	var user1 *User                    // 声明结构体指针
	fmt.Printf("user1 = %#v\n", user1) // 未初始化的时候指针为nil

	var user2 *User = &User{}

	// 指针类型也可以直接通过.成员变量名来访问  ==> Go语言的语法糖实现了指针的自动转换，达到简化语法的目的
	// user2.Age = 18
	(*user2).Age = 18
	user2.AvataUrl = "https://img2020.cnblogs.com/blog/1188507/202111/1188507-20211116110743641-1021429203.png"
	user2.Sex = 1
	user2.Username = "Tom"

	fmt.Printf("user2 = %#v\n", user2)

	user3 := &User{Username: "Bob", Sex: 1, Age: 19, AvataUrl: "https://baidu.com"}
	fmt.Printf("user3 = %#v\n", user3)

	user4 := new(User) // new分配内存地址，返回的是指针
	user4.Age = 21
	user4.Sex = 0
	user4.Username = "Lucy"
	user4.AvataUrl = "https://sohu.com"

	fmt.Printf("user4 = %#v\n", user4)
}
