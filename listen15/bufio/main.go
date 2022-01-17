package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func bufio_read_file() {
	if file, err := os.Open("test.txt"); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
			} else {
				fmt.Println(strings.TrimSpace(line))
			}
		}
	}
}

func main() {
	bufio_read_file()
}