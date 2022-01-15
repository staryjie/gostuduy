package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Student struct {
	Age   int
	Name  string
	Score float64
}

func storeStudents(n int) {
	rand.Seed(time.Now().UnixNano())
	stuMap := make(map[int]Student, n)

	for i := 11; i <= n+10; i++ {
		id := i
		stu := Student{Age: rand.Intn(15) + 10, Name: fmt.Sprintf("stu%02d", i), Score: float64(rand.Intn(60)+40) + rand.Float64()}
		stuMap[id] = stu
	}

	keys := make([]int, 0, n)
	for k, _ := range stuMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("id=%d's student info\tName: %s\tAge: %d\tScore: %.2f\n", k, stuMap[k].Name, stuMap[k].Age, stuMap[k].Score)
	}
}

func main() {
	storeStudents(10)
}
