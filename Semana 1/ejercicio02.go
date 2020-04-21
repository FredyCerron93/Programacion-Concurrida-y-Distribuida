package main

import "fmt"

func main() {
	//declaración de variables
	var largo float64 = 12.36
	ancho := 19

	fmt.Println("El area calculada es: ", largo*float64(ancho))
	fmt.Println("Es el largo mayor al ancho: ", largo > float64(ancho))

}
