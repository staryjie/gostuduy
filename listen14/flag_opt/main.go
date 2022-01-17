package main

import (
	"flag"
	"fmt"
	"strings"
)

// flag库Go官方提供的用来处理命令行传入的参数，os.Args也可以处理，但是并不是很理想
// 第三方提供的库有 cobra cli等

// 自定义参数解析
type Custom []string

// 实现String()方法
func (c *Custom) String() string {
	return fmt.Sprintf("%v", *c)
}

// 实现Set方法，Set方法决定了如何解析flag的值
func (c *Custom) Set(s string) error {
	for _, v := range strings.Split(s, ",") {
		*c = append(*c, v)
	}
	return nil
}

// 声明变量用于接收命令行传入的参数值
var (
	name    string
	age     int
	address *string // 指针
	id      *int    // 指针
	cus     Custom
)

// init()函数中进行传参处理
func init() {
	// 通过传入变量地址的方式，绑定命令行参数到string变量
	flag.StringVar(&name, // 存放值的参数的指针，传入指针
		"name", // 命令行参数的名称
		"匿名",   // 命令行参数没有输入时的默认值
		"您的姓名") // 参数的描述信息,help时会显示
	flag.IntVar(&age, "age", -1, "您的年龄")

	// 直接返回参数值的指针
	address = flag.String("address", "未知", "您的住址")
	id = flag.Int("id", -1, "您的Id号")
	flag.Var(&cus, "custom", "输入随意的内容")
}

func main() {
	// 传参处理
	flag.Parse()

	// 传参处理之后把值传递给对应的变量，可以使用
	fmt.Printf("%s，您好，您的编号为%d，您的年龄为%d，您的住址为%s。\n", name, *id, age, *address)

	fmt.Println("遍历所有的输入参数(开始)")
	flag.Visit(func(f *flag.Flag) { // Visit函数只会遍历命令行中输入的参数
		fmt.Printf("参数名[%s], 参数值[%s], 默认值[%s], 描述信息[%s]\n", f.Name, f.Value, f.DefValue, f.Usage)
	})
	fmt.Println("遍历所有的输入参数(结束)")

	fmt.Println("遍历所有的参数(开始)")
	flag.VisitAll(func(f *flag.Flag) { // VisitAll函数会遍历所有定义的参数，无论是否在命令行中输入
		fmt.Printf("参数名[%s], 参数值[%s], 默认值[%s], 描述信息[%s]\n", f.Name, f.Value, f.DefValue, f.Usage)
	})
	fmt.Println("遍历所有的参数(结束)")

	fmt.Printf("%v\n", cus)
}

// 必须要先build然后执行可执行文件
// 必须在命令行执行可执行文件
// go build main.go
// ./main --help
// ./main -name xxx -age xx -id xx -address xxx
