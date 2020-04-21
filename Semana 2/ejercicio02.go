package main

import (
	"fmt"
	"time"
)

var n int

func p() {
	var temp int
	temp = n
	n = temp + 1
	time.Sleep(time.Second)
}

func q() {
	var temp int
	temp = n
	n = temp + 1
	time.Sleep(time.Second)
}
func main() {
	n = 0
	go p()
	go q()
	time.Sleep(time.Second)
	fmt.Println(n)
}
