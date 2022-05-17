package main

import (
	"C"
	"time"
)

var c chan int

func decrement(n int) {
	for n > 0 {
		n -= 1
	}
}

//export count_time
func count_time() *C.char {
	start := time.Now()
	decrement(100000000)
	total_time := time.Since(start).String()
	return C.CString(total_time)
}

func main() {

}
