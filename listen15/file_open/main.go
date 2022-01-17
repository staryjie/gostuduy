package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// file.Read()
func read_byte() {
	// 以字节的方式读取指定字节的内容
	file, err := os.Open("test.txt") // 只读方式打开
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	bs := make([]byte, 1024)
	n, _ := file.Read(bs)
	fmt.Printf("读取了%d字节, 内容为\n%s\n", n, string(bs[:n]))
}

func read_file_byte() {
	// 以字节的方式读取整个文件
	if file, err := os.Open("test.txt"); err != nil { // 只读方式打开文件
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

// file.ReadAt()
func read_byte_at() {
	// 以字节方式读取文件，从指定的位置开始读取
	if file, err := os.Open("test.txt"); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		bs := make([]byte, 1024)

		n, _ := file.ReadAt(bs, 3)
		fmt.Printf("读取了%d字节, 内容为\n%s\n", n, string(bs[:n]))
	}
}

func main() {
	//read_byte()
	//read_file_byte()
	read_byte_at()
}
