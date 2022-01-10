package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
		D
		E = 8
		F
		G = iota
		H
	)

	const (
		A1 = iota
		A2
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)

	fmt.Println("--------")

	fmt.Println(A1)
	fmt.Println(A2)
}
