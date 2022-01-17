package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		a int
		s string
		f float64
	)
	//fmt.Fscanf(os.Stdin, "%d %s %f", &a, &s, &f)
	//fmt.Fscan(os.Stdin, &a, &s, &f)
	fmt.Fscanln(os.Stdin, &a, &s, &f)

	//fmt.Fprintf(os.Stdout, "a = %d  s = %s  f = %.2f\n", a, s, f)
	//fmt.Fprint(os.Stdout, a, s, f)
	fmt.Fprintln(os.Stdout, a, s, f)
}