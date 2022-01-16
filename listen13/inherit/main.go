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
	*BuruAnimal
	Feet string
}

func (d *Dog) Eat() {
	fmt.Println("Dog Eat.")
}

/*
func (d *Dog) Talk() { // 重写父类方法
	fmt.Println("wang wang!")
}
*/

type BuruAnimal struct {
	*Animal
}

func (b *BuruAnimal) Talk() {
	fmt.Printf("Hi, I'm %s\n", b.Name)
}

func main() {
	var dog *Dog = &Dog{Feet: "four legs", Animal: &Animal{Name: "小白", Sex: "雄性"}}
	// dog.Talk()  // ambiguous selector dog.Talk
	dog.Animal.Talk()
	dog.Eat()
}
