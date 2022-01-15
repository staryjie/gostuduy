package main

import (
	"fmt"
	"math/rand"
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

	for i := 1; i <= n; i++ {
		id := i
		stu := Student{Age: rand.Intn(15) + 10, Name: fmt.Sprintf("stu%02d", i), Score: float64(rand.Intn(60) + 40 + int(rand.Float64()))}
		stuMap[id] = stu
	}

	for k, v := range stuMap {
		fmt.Printf("id=%d's student info Name: %s\tAge: %d\tScore:%.2f\n", k, v.Name, v.Age, v.Score)
	}
}

func main() {
	storeStudents(10)
}
