package main

import (
	"fmt"
	"strings"
)

func statWordCount(str string) map[string]int {
	words := strings.Split(str, " ")

	result := make(map[string]int, len(words))

	for _, word := range words {
		/*
			count, ok := result[word]
			if !ok {
				result[word] = 1
			} else {
				result[word] = count + 1
			}
		*/
		result[word] += 1
	}
	return result
}

func main() {
	str := "how do you do ? are you kidding"
	result := statWordCount(str)
	fmt.Printf("result: %#v\n", result)
}
