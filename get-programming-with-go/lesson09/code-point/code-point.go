package main

import "fmt"

func main() {
	// 为了打印出字符而不是数字值本身，我们可以在Printf中使用格式化变量%c而不是%v
	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)

	smile := '😃'
	fmt.Printf("%c %[1]v\n", smile)

	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)
}
