package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	if src, err := os.Open(srcName); err != nil {
		//fmt.Println(err)
		return -1, err
	} else {
		defer src.Close()
		if dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
			//fmt.Println(err)
			return -1, err
		} else {
			defer dst.Close()
			return io.Copy(dst, src)
		}
	}
	return 
}

func main() {
	if _, err := CopyFile("target.txt", "main.go"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Copy Done!")
	}
}
