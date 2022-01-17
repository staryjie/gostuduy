package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if file, err := os.Open("test.txt"); err != nil {  // 只读方式打开文件
		fmt.Println(err)
	} else {
		defer file.Close() // 关闭文件
		var buffer strings.Builder
		for {
			bs := make([]byte, 1024) // 指定读取文件的字节数
			if n, err := file.Read(bs); err != nil {
				if err == io.EOF { // 读到文件末尾
					break
				}
				fmt.Println(err)
			} else {
				fmt.Printf("从文件中读取了%d字节。\n", n)
				buffer.WriteString(string(bs)) // 写入buffer
			}
		}
		fmt.Println("文件内容如下: \n----------------------------")
		fmt.Println(buffer.String())
		fmt.Println("----------------------------")
	}
}
