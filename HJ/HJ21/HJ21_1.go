package main

import "fmt"

func main() {
	var in string
	fmt.Scanf("%s", &in)
	tmp := 'a' - 'A' + 1
	a := []int{2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9}
	for _, v := range in {
		if v >= '0' && v <= '9' {
			fmt.Print(string(v))
		} else if v >= 'A' && v <= 'Z' {
			s := v + tmp
			if s > 'z' {
				s = 'a'
			}
			fmt.Print(string(s))
		} else {
			fmt.Print(a[v-'a'])
		}
	}
}
