package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string  `json:"username"`
	Sex      int     `json:"Gender"`
	Score    float32 `json:"score"`
	Age      int32   `json:"age"`
}

func main() {
	user := &User{
		UserName: "user01",
		Sex:      1,
		Score:    99.2,
		Age:      18,
	}

	fmt.Printf("user = %#v\n", user)

	data, _ := json.Marshal(user)
	fmt.Printf("user json = %s\n", string(data))
}
