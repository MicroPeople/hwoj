package main

import (
	"fmt"
)

func main() {
	var str string
	count := 0
	fmt.Scanf("%s", &str)

	strMap := make(map[byte]int)

	for i := range str {
		if _, ok := strMap[str[i]]; !ok {
			count++
			strMap[str[i]] = 1
		}
	}
	fmt.Println(count)
}
