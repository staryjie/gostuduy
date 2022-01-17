package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if context, err := ioutil.ReadFile("./main.go"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(context))
	}
	
}
