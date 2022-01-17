package main

import "fmt"

type Animaler interface {
	Talk()
	Eat()
	Name() string
}

type BuruAnimal interface {
	Production() string
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

func (d Dog) Production() string {
	production := "viviparous"
	return production
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

func (p Pig) Production() string {
	production := "viviparous"
	return production
}

func (p Pig) Name() string {
	return p.name
}


func main() {
	d := Dog{name: "小白"}
	var a Animaler
	var b BuruAnimal
	a = d
	b = d
	a.Talk()
	fmt.Println(b.Production())
}