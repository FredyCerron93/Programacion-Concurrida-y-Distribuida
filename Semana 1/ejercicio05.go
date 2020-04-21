package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	bufferIn := bufio.NewReader(os.Stdin)

	menu :=
		`
		Bienvenido
		[1] Pizza
		[2] Empanadas
		¿Qué prefieres?
	`
	fmt.Println(menu)

	ingreso, _ := bufferIn.ReadString('\n')
	dato := strings.TrimRight(ingreso, "\r\n")

	switch dato {
	case "1":
		fmt.Println("Elegiste Pizza")
	case "2":
		fmt.Println("Elegiste Empanadas")
	default:
		fmt.Println("Opción no válida")
	}

}
