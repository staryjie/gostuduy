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

func init_switch() {
	switch a := 3; a {
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

func case_mutil() {
	var a string = "i"
	switch a {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

func switch_expr() {
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}
}

func switch_fallthrough() {
	num := 50
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("0 <= num <= 50")
		fallthrough
	case num > 50 && num <= 100:
		fmt.Println("50 < num <= 100")
	case num > 100 && num <= 150:
		fmt.Println("100 < num <= 150")
	}
}

func main() {
	// switch_demo()
	// init_switch()
	// case_mutil()
	// switch_expr()
	switch_fallthrough()
}
