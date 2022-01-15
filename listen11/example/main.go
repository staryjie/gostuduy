package main

import "fmt"

var (
	coins = 100
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func calcCoin(username string) int {
	var castCoins int // 统计分出去多少金币
	// 遍历username中的每个字符，判断该分多少金币
	for _, char := range username {
		switch char {
		case 'a', 'A', 'e', 'E':
			castCoins += 1
		case 'i', 'I':
			castCoins += 2
		case 'o', 'O':
			castCoins += 3
		case 'u', 'U':
			castCoins += 5
		}
	}

	return castCoins
}

func coinsLeft() int {
	var left int = coins
	// 遍历用户名切片
	for _, username := range users {
		castCoins := calcCoin(username)
		left -= castCoins
		if _, ok := distribution[username]; ok {
			distribution[username] += 1
		} else {
			distribution[username] = castCoins
		}
	}

	return left
}

func main() {
	left := coinsLeft()
	for username, coin := range distribution {
		fmt.Printf("user %s have %d coins\n", username, coin)
	}

	fmt.Printf("There are %d coins left\n", left)
}
