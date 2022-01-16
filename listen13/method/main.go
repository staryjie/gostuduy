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

func main() {
	p1 := People{Name: "Tom", Country: "America"}
	p1.PeopleInfo()
}
