package main

import "fmt"

func main() {
	//declarar un arreglo
	arreglo := [7]int{0, 4, 1, 6, 10, 9, 7}

	//estructuras repetitivas
	for i, j := range arreglo {
		fmt.Printf("Valor de j es %d en la vuelta #%d\n", j, i)
	}

	for _, j := range arreglo {
		fmt.Printf("El valor de j es %d \n", j)
	}

	//funciona como while(1) en c++
	i := 0
	for {
		if i == 5 {
			break
		}
		i++
	}
	println("Intentos=", i)

	j := 0
	for i > 1 {
		if j == 5 {
			break
		}
		j++
	}
	println("Intentos=", j)
}
