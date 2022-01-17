package main

import "fmt"

type Student struct {
	Name string
	Sex int
}

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

func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("%T = %d\n", i, i)
	case string:
		fmt.Printf("%T = %s\n", i, i)
	case float32:
		fmt.Printf("%T = %f\n", i, i)
	case float64:
		fmt.Printf("%T = %f\n", i, i)
	default:
		fmt.Printf("%T = %#v\n", i, i)
	}
}

func main() {
	i := 12
	getType(i)
	findType(i)

	s := "hello"
	getType(s)
	findType(s)

	var f64 float64 = 3.14
	getType(f64)
	findType(f64)

	var f32 float32 = 4.17
	getType(f32)
	findType(f32)

	stu := Student{
		Name: "Tom",
		Sex: 1,
	}
	findType(stu)
}