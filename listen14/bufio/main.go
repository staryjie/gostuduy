package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string

	// fmt包无法实现终端读取一行数据
	//fmt.Scanf("%s", &s)
	//fmt.Println(s)
	reader := bufio.NewReader(os.Stdin) // 从终端读取
	str, _ = reader.ReadString('\n')
	//fmt.Printf("You enter: %s\n", str)
	fmt.Fprintf(os.Stdout, "You enter: " + str)
}
