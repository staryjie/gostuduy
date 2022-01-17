package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("args[0] = %#v\n", os.Args[0])
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			if i == 0 {  // 程序名不输出
				continue
			}
			fmt.Printf("arfgs[%d] = %v\n", i, arg)
		}
	}
}
