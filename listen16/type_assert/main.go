package main

import "fmt"

func getType(i interface{}) {
	if t, ok := i.(int); ok {
		fmt.Printf("%T = %d\n", t, t)
		return
	}

	if t, ok := i.(string); ok {
		fmt.Printf("%T = %s\n", t, t)
		return
	}

	if t, ok := i.(float64); ok {
		fmt.Printf("%T = %f\n", t, t)
		return
	}
	if t, ok := i.(float32); ok {
		fmt.Printf("%T = %f\n", t, t)
		return
	}
}

func main() {
	i := 12
	getType(i)
	
	s := "hello"
	getType(s)

	var f64 float64 = 3.14
	getType(f64)

	var f32 float32 = 4.17
	getType(f32)
}