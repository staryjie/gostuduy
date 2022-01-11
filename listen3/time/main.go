package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testTime() {
	now := time.Now()
	fmt.Printf("%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 时
	minute := now.Minute() // 分
	second := now.Second() // 秒

	fmt.Printf("当前时间: %d-%d-%2d %d:%d:%d\n", year, int(month), day, hour, minute, second)

	// 时间戳
	fmt.Println("当前时间戳:", now.Unix())

	// 获取时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println("时区:", loc)

}

// 时间戳转换
func TranslateTimeStamp(timestamp int64) string {
	timeObj := time.Unix(timestamp, 0)

	res := fmt.Sprintf("%d ==> %d-%d-%2d %d:%d:%d\n",
		timestamp, timeObj.Year(), int(timeObj.Month()), timeObj.Day(),
		timeObj.Hour(), timeObj.Minute(), timeObj.Second())

	return res
}

func ProcessTask() {
	fmt.Println("Do Task!")
}

func timeTicker() {
	ticker := time.Tick(time.Second) // 返回一个channel
	for i := range ticker {          // 遍历channel实现定时器
		fmt.Println(i.Format("2006-01-02 15:04:05"))
		ProcessTask()
	}
}

func timeConst() {
	// 时间常量
	fmt.Printf("纳秒: %d\n", time.Nanosecond)
	fmt.Printf("微秒: %d\n", time.Millisecond)
	fmt.Printf("毫秒: %d\n", time.Microsecond)
	fmt.Printf("秒  : %d\n", time.Second)
}

func timeFormat() {
	now := time.Now()

	fmt.Println(now.Format("2006/01/02 15:04:05"))
}

func timeConsume() {
	begin := time.Now()
	fmt.Println("Do some tasks...")
	rand.Seed(time.Now().UnixMilli())
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	res := time.Since(begin).Microseconds()

	fmt.Printf("程序执行耗时: %d微秒\n", res)
}

func main() {
	// testTime()

	// 时间戳转换
	/*
		timestamp := time.Now().Unix()
		result := TranslateTimeStamp(timestamp)
		fmt.Println(result)
	*/

	// 定时器
	// timeTicker()

	// 时间常量
	// timeConst()
	// timeFormat()

	// 统计时间差
	timeConsume()
}
