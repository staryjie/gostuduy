package main

import (
	"os"
)

/*
	终端的底层也是文件
	var (
		Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
		Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	)
 */

func main() {
	var buf [16]byte
	os.Stdin.Read(buf[:])
	
	//fmt.Println(string(buf[:]))
	os.Stdout.WriteString("You enter: "+ string(buf[:]))
}

