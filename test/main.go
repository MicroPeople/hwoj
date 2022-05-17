package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(" ", i)
	}
	var a, b int64 = 4, 5
	var c int64 = (a + b) / 2
	println(c)
}
