package main

import "fmt"

//strings.Repeat

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0
	//for i := 0; i < len(cipherText); i++ {
	//	// A=0, B=1, ... Z=25
	//	c := cipherText[i] - 'A'
	//	k := keyword[keyIndex] - 'A'
	//	// 加密字母-关键字字母
	//	c = (c-k+26)%26 + 'A'
	//	message += string(c)
	//	// 对keyIndex执行自增操作
	//	keyIndex++
	//	keyIndex %= len(keyword)
	//}

	for i := range cipherText {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		fmt.Printf("%v ", c)
		k := keyword[keyIndex] - 'A'
		// 加密字母-关键字字母
		c = (c-k+26)%26 + 'A'
		message += string(c)
		// 对keyIndex执行自增操作
		keyIndex++
		keyIndex %= len(keyword)
	}

	fmt.Println(message)
	message = ""
	keyIndex = 0

	for _, c := range cipherText {
		// A=0, B=1, ... Z=25
		//c := cipherText[i] - 'A'
		c1 := uint8(c) - 'A'
		fmt.Printf("%v ", c1)
		k := keyword[keyIndex] - 'A'
		// 加密字母-关键字字母
		c1 = (c1-k+26)%26 + 'A'
		message += string(c1)
		// 对keyIndex执行自增操作
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)
}
