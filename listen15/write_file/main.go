package main

import (
	"bufio"
	"fmt"
	"os"
)
/*
os.O_CREATE    创建文件
os.O_APPEND    追加文件
os.O_RDONLY    只读
os.O_WRONLY    只写
os.O_RDWR      读写
os.O_TRUNC     清空文件内容
 */

func write_string(filename string) {
	if file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		str := "Hello World!"
		file.WriteString(str)
	}
}

func write_bytes(filename string) {
	if file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		str := "Hello World!"
		file.Write([]byte(str))
	}
}

func bufio_write_string(filename string) {
	if file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		writer := bufio.NewWriter(file)
		str := "Hello World!"
		writer.WriteString(str)
		writer.Flush()  // 将buf中的内容刷到磁盘上
	}
}

func bufio_write_bytes(filename string) {
	if file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		writer := bufio.NewWriter(file)
		str := "Hello World!"
		writer.Write([]byte(str))
		writer.Flush()  // 将buf中的内容刷到磁盘上
	}
}

func main() {
	filename := "./listen15/write_file/test.txt"
	//write_string(filename)
	//write_bytes(filename)
	//bufio_write_string(filename)
	bufio_write_bytes(filename)
}
