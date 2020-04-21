package main

import "fmt"

var wantp bool = false
var wantq bool = false

func p() {
	for {
		fmt.Println("P, Seccion NC Linea 1")
		fmt.Println("P, Seccion NC Linea 2")
		wantp = true
		for wantq {
			//espera
		}
		fmt.Println("P, Seccion C Linea 1")
		fmt.Println("P, Seccion C Linea 2")
		wantp = false
	}
}

func q() {
	for {
		fmt.Println("Q, Seccion NC Linea 1")
		fmt.Println("Q, Seccion NC Linea 2")
		wantq = true
		for wantp {
			//espera
		}
		fmt.Println("Q, Seccion C Linea 1")
		fmt.Println("Q, Seccion C Linea 2")
		wantq = false
	}
}

func main() {
	go p()
	q()
}
