package main

import (
	"fmt"
)

type People struct {
	Name    string
	Country string
}

func (p People) PeopleInfo() {
	fmt.Printf("%s is come from %s!\n", p.Name, p.Country)
}

func (p *People) Set(name, country string) {
	p.Name = name
	p.Country = country
}

func main() {
	p1 := People{Name: "Tom", Country: "America"}
	p1.PeopleInfo()
	(&p1).Set("Jack", "Japan")
	p1.PeopleInfo()
	p1.Set("Haha", "There") // 实现指针类型的方法，值类型也可以直接调用，这是Go语言语法糖实现的
	p1.PeopleInfo()

	p2 := &People{Name: "Bob", Country: "America"}
	p2.Set("Tim", "England")
	p2.PeopleInfo()
}

/*
	值类型和指针类型的使用场景：
	1. 需要修改实例中成员的值的时候，需要传递指针类型
	2. 接受者是大对虾(成员变量较多)，传递值类型拷贝大家比较大
	3. 通常情况下推荐使用指针类型作为参数接收者
*/
