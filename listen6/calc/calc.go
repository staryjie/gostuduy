package calc

var a int = 100
var A int = 200 // 可导出

func Add(a, b int) int { // 可导出
	return a + b
}

func sub(a, b int) int {
	return a - b
}
