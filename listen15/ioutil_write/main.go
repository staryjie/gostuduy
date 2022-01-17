package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "Hello World!"

	if err := ioutil.WriteFile("./listen15/ioutil_write/test.txt", []byte(str), 0755); err != nil {
		fmt.Println(err)
		return
	}
}
