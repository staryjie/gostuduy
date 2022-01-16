package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Sex  string
}

type Class struct {
	Name     string
	Count    int
	Students []*Student
}

func main() {
	c := &Class{
		Name:  "101",
		Count: 10,
	}

	for i := 1; i <= 10; i++ {
		stu := &Student{
			Id:   i,
			Name: fmt.Sprintf("stu%d", i),
			Sex:  "male",
		}
		c.Students = append(c.Students, stu)
	}

	if data, err := json.Marshal(c); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
