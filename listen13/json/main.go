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

var rawJson = `{"Name":"101","Count":10,"Students":[{"Id":1,"Name":"stu1","Sex":"male"},{"Id":2,"Name":"stu2","Sex":"male"},{"Id":3,"Name":"stu3","Sex":"male"},{"Id":4,"Name":"stu4","Sex":"male"},{"Id":5,"Name":"stu5","Sex":"male"},{"Id":6,"Name":"stu6","Sex":"male"},{"Id":7,"Name":"stu7","Sex":"male"},{"Id":8,"Name":"stu8","Sex":"male"},{"Id":9,"Name":"stu9","Sex":"male"},{"Id":10,"Name":"stu10","Sex":"male"}]}`

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

	// json反序列换
	var c1 = &Class{}
	if err := json.Unmarshal([]byte(rawJson), c1); err != nil { // 结构体实例必须是指针类型
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", *c1)
	}

}
