package main

import "fmt"

type User struct { // 结构体的定义
	Username string
	Sex      int
	Age      int
	AvataUrl string
}

func main() {
	var user User // 声明结构体实例
	// 对结构体实例成员属性赋值
	user.Age = 18
	user.AvataUrl = "https://img2020.cnblogs.com/blog/1188507/202111/1188507-20211116110743641-1021429203.png"
	user.Sex = 1
	user.Username = "Tom"

	// 结构体实例成员属性访问
	fmt.Printf("Username:%s age:%d sex:%d avataurl:%s\n", user.Username, user.Age, user.Sex, user.AvataUrl)

	user2 := User{Username: "Tim", Age: 20, Sex: 1, AvataUrl: "https://www.baidu.com"}
	fmt.Printf("Username:%s age:%d sex:%d avataurl:%s\n", user2.Username, user2.Age, user2.Sex, user2.AvataUrl)
}
