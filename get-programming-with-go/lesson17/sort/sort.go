package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	// planets变量将从[]string类型转换为在sort包中声明的StringSlice类型。
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}
