package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0

	fmt.Scan(&n)
	li := make([]string, 0)
	for i := 0; i < n; i++ {

		s := ""
		fmt.Scan(&s)
		li = append(li, s)
	}
	sort.Strings(li)
	for _, v := range li {
		fmt.Println(v)
	}
}
