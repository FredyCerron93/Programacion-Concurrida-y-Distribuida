package main

import (
	"fmt"
	"time"
)

const FINAL = 100 * time.Millisecond

func saludo(i int) {
	time.Sleep(10 * time.Duration(i%5) * time.Millisecond)
	fmt.Println("Bienvenido al curso de Prod.Concurrente ", i)
}

func main() {
	for i := 1; i <= 6; i++ {
		//lanzar concurrentemente
		go saludo(i)
	}

	time.Sleep(FINAL)
}
