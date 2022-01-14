package main

import (
	"fmt"
)

func printAddr() {
	var a int = 30
	fmt.Printf("Variable a's address: %p\n", &a)
}

func modifyInt(a *int) {
	*a = *a + 100
}

func testModifyInt() {
	var a int = 100
	modifyInt(&a)
	fmt.Println(a)
}

func main() {
	printAddr()
	testModifyInt()
}
