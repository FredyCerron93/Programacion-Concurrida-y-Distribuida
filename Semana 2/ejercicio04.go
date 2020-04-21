package main

import (
	"fmt"
	"time"
)

var n int

func p() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		time.Sleep(10 * time.Duration(i%5) * time.Millisecond)
		n = temp + 1
		time.Sleep(10 * time.Duration(i%5) * time.Millisecond)
	}
}

func main() {
	go p()
	go p()
	time.Sleep(time.Second)
	fmt.Println(n)
}
