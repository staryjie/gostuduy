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

	p2 := &People{Name: "Bob", Country: "America"}
	p2.Set("Tim", "England")
	p2.PeopleInfo()
}
