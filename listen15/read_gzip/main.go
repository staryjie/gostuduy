package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	if file, err := os.Open("./main.gz"); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		if fz, err := gzip.NewReader(file); err != nil {
			fmt.Println(err)
		} else {
			bufreader := bufio.NewReader(fz)
			var context []byte
			bs := make([]byte, 1024)
			for {
				if n, err := bufreader.Read(bs); err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println(err)
				} else {
					context = append(context, bs[:n]...)
				}
			}
			fmt.Println(string(context))
		}
	}
}
