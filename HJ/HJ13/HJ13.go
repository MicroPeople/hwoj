package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	res := strings.Split(input.Text(), " ")
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i] + " ")
	}

}
