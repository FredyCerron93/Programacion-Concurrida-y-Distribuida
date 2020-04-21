package main

import(
	"fmt"
	"math/rand"
	"time"
)

func imprimir(ch chan string, texto string) {
	//pasar un valor texto mediante el canal
	ch <- texto

	//tiempo de espera
	time.Sleep(time.Duration(rand.Int()) * time.Nanosecond)
}

func main(){
	//Crear un canal de tipo string
	ch := make(chan string)
	//pasando 
	go imprimir(ch,"enviar 1")
	go imprimir(ch,"enviar 2")
	 // ciclo infinito, vamos a imprimir la informacion que viene en el canal
	 fo{
		 fmt.Println(<-ch)
	 }
}
