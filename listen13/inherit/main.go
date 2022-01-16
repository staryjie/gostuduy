package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Sex  string
}

func (a *Animal) Talk() {
	fmt.Printf("I'm %s\n", a.Name)
}

type Dog struct {
	// Animal
	*Animal // 实现"继承"
	Feet    string
}

func (d *Dog) Eat() {
	fmt.Println("Dog Eat.")
}

func (d *Dog) Talk() { // 重写父类方法
	fmt.Println("wang wang!")
}

func main() {
	var dog *Dog = &Dog{Feet: "four legs", Animal: &Animal{Name: "小白", Sex: "雄性"}}
	dog.Talk()
	dog.Eat()
}
