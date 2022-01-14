package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var m map[string]int = make(map[string]int, 1024)

	for i := 1; i <= 128; i++ {
		key := fmt.Sprintf("stu%2d", i)
		value := rand.Intn(1000)
		m[key] = value
	}

	for k, v := range m {
		fmt.Printf("m[%s] = %d\n", k, v)
	}
}
