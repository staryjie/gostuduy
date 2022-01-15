package main

import (
	"fmt"
)

type Address struct {
	Country  string
	Province string
	City     string
}

type User struct {
	Username string
	Age      int
	Sex      int
	City     string
	int      // 匿名字段直接使用数据类型作为变量
	string
	// Address // 结构体嵌套,直接使用结构体名做为匿名成员变量
	*Address
}

func main() {
	user := &User{
		Username: "Tom",
		Age:      20,
		Sex:      1,
		City:     "Beijing",
		int:      12,
		string:   "Hi",
		// Address:  Address{Country: "China", Province: "ZHejiang", City: "Hangzhou"},
		Address: &Address{Country: "China", Province: "ZHejiang", City: "Hangzhou"},
	}
	fmt.Printf("user = %#v\n", user)
	fmt.Printf("user.Address = %#v\n", user.Address)

	fmt.Println(user.Country)
	fmt.Println(user.Province)
	fmt.Println(user.City)
	fmt.Println(user.Address.City)
}
