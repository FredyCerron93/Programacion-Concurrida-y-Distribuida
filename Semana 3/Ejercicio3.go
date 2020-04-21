package main

import (
	"fmt"
)

var turn int = 1

func p() {
	for {
		fmt.Println("P, seccion NC l1")
		fmt.Println("P, seccion NC l2")
		for turn != 1 {
			//Espereando
		}
		fmt.Println("P, seccion C l1")
		fmt.Println("P, seccion C l2")

		turn = 2
	}
}

func q() {
	for {
		fmt.Println("Q, seccion NC l1")
		fmt.Println("Q, seccion NC l2")

		for turn != 2 {
			// Esperando
		}

		fmt.Println("Q, seccion C l1")
		fmt.Println("Q, seccion C l2")

		turn = 1
	}
}

func main() {
	go p()
	q()
}
