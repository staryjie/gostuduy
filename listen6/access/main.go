package main

import (
	"fmt"

	"github.com/staryjie/gostudy/listen6/calc"
)

func main() {
	var s1 int = 100
	var s2 int = 300
	sum := calc.Add(s1, s2)
	fmt.Println(sum)

	// sub := calc.sub(s2, s1)  // sub小写开头，无法在包外访问

	fmt.Printf("N1 = %d\n", calc.Num)
	// fmt.Printf("a = %d\n", calc.a)  // a小写无法导出，无法在包外访问
	fmt.Printf("n1 = %d\n", calc.Getnum())
}
