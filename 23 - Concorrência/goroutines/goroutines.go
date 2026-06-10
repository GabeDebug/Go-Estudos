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

		O que é uma Goroutine?
		Uma goroutine é uma função que é executada concorrentemente (em paralelo ou de forma intercalada) com outras goroutines. Elas são semelhantes a threads, mas muito mais leves.

		Uma goroutine usa apenas alguns kilobytes de memória na pilha (stack).
		O runtime do Go gerencia automaticamente as goroutines (não o sistema operacional).
		Você pode ter dezenas de milhares ou até centenas de milhares de goroutines rodando ao mesmo tempo sem problemas.
	*/

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
