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
