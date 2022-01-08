package main

import (
	"fmt"
	"time"
)

func calc() {
	for i := 0; i < 10; i++ {
		fmt.Printf("run %d times\n", i+1)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Calc finished!")
}

func main() {
	// calc() // 直接调用，会阻塞主进程
	go calc() // 高并发，启动一个线程去执行 calc

	fmt.Println("i exited!")
	time.Sleep(15 * time.Second) // 休眠，让子线程能执行完成
}
