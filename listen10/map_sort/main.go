package main

import (
	"fmt"
	"math/rand"
	"sort"
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

	var keys []string = make([]string, 0, 128)

	for k, _ := range m {
		// fmt.Printf("m[%s] = %d\n", k, v)
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("key:%s value:%d\n", key, m[key])
	}
}
