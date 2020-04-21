package main

import "fmt"

func sumatoria(tope float64) float64 {
	if tope == 1 {
		return tope
	} else {
		return tope + sumatoria(tope-1)
	}
}

func main() {
	fmt.Println("Ingrese un número: ")
	var input float64
	fmt.Scanf("%f", &input)

	//fmt.Println(input * 2)

	fmt.Println("Sumatoria de número")
	fmt.Printf("La sumatoria hasta %f es = %f", input, sumatoria(input))
}
