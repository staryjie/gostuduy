package main

import "fmt"

type Animaler interface {
	Talk()
	Eat()
	Name() string
}

type Describe interface {
	Describe() string
}

type AdvanceAnimal interface {
	Animaler
	Describe
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

func (d Dog) Describe() string {
	msg := "This is a dog"
	return msg
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

func (p Pig) Describe() string {
	msg := "This is a pig"
	return msg
}

func (p Pig) Name() string {
	return p.name
}

func main() {
	d := Dog{name: "小白"}
	var a AdvanceAnimal
	a = d
	a.Eat()
	fmt.Println(a.Describe())
}