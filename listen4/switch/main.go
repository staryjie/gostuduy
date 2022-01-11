package main

import "fmt"

func switch_demo() {
	var a int = 2
	switch a {
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	case 3:
		fmt.Println("a = 3")
	case 4:
		fmt.Println("a = 4")
	case 5:
		fmt.Println("a = 5")
	default:
		fmt.Printf("a = %d\n", a)
	}
}

func main() {
	switch_demo()
}
