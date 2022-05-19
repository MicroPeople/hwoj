package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	H := make(map[string]string)
	fmt.Scanf("%s\n", &str)
	H["abc"] = "2"
	H["def"] = "3"
	H["ghi"] = "4"
	H["jkl"] = "5"
	H["mno"] = "6"
	H["pqrs"] = "7"
	H["tuv"] = "8"
	H["wxyz"] = "9"
	for _, v := range str {
		for k, v2 := range H {
			if strings.Contains(k, string(v)) {
				fmt.Printf("%s", v2)
			}
		}
		if v >= 'A' && v < 'Z' {
			var str1 string = strings.ToLower(string(v + 1))
			fmt.Printf("%s", str1)
		} else if v == 'Z' {
			fmt.Printf("%s", "a")
		} else if v >= '0' && v <= '9' {
			fmt.Printf("%s", string(v))
		}
	}
}
