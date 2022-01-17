package main

import "fmt"

type Animaler interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {
	name string
}

func (d Dog) Talk() {
	fmt.Printf("%s: 汪汪汪\n", d.name)
}

func (d Dog) Eat() {
	fmt.Printf("%s在啃骨头\n", d.name)
}

func (d Dog) Name() string {
	return d.name
}

type Pig struct {
	name string
}

func (p Pig) Talk() {
	fmt.Printf("%s: 杠杠杠\n", p.name)
}

func (p Pig) Eat() {
	fmt.Printf("%s在吃草\n", p.name)
}

func (p Pig) Name() string {
	return p.name
}

func main() {
	dog := Dog{name: "小白"}
	dog.Talk()
	dog.Eat()
	fmt.Println(dog.Name())

	var a Animaler
	a = dog
	a.Talk()
	a.Eat()
	fmt.Println(a.Name())

	fmt.Println("----------")

	pig := Pig{name: "小花"}
	pig.Talk()
	pig.Eat()
	fmt.Println(pig.Name())

	fmt.Println("----------")

	// 指针类型
	d := &Dog{name: "小黑"}
	d.Eat()
	d.Talk()
	fmt.Println(d.Name())
}