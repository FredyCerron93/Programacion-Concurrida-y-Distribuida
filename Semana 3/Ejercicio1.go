package main

import "fmt"

var ch chan bool //declaracion de un canal

func p() {
	fmt.Println("Envia dato")
	ch <- true
}

func main() {
	ch = make(chan bool) //nos permite crear un canal de tipo booleano y guardarlo en la variable
	//ch <- true
	go p()
	<-ch

	fmt.Println("recibe")
}
