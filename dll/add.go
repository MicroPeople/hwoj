package main

import "C"

import "fmt"

type name struct {
}

//export add
func add(a, b int) int {
	var sum int = a + b
	fmt.Println(sum)
	return sum
}

func main() {

}
