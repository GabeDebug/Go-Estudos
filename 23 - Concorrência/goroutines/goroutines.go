package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrencia != Paralelismo
	go escrever("Olá mundo") // um jeito para contorna isso e usando uma goroutines
	go escrever("programando em Golang")
	escrever("gato")
	/*
		quando usamos uma goroutines e dizer que pode executa essa função
		mais não deixa ela termina pra executa a proxima
	*/

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
