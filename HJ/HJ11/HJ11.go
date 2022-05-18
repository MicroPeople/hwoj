package main

import "fmt"

//func main() {
//	var num int16
//	fmt.Scanf("%d", &num)
//	for num >= 10 {
//		fmt.Printf("%d", num%10)
//		num /= 10
//	}
//	fmt.Printf("%d", num)
//}
//
func main() {
	var s string
	fmt.Scanf("%s", &s)
	//for i, ch := range str {
	//	fmt.Printf("str[%d]=%c\n", i, ch)
	//}
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	fmt.Printf("%s", string(r))
}
