package main

import "fmt"

func main() {
	/*
	 O defer no Golang (Go)
	 é uma palavra-chave usada para adiar a execução de uma
	 função até o final da função atual.
	*/
	fmt.Println("Inicio")

	defer fmt.Println("Isso roda por ultimo")
	// usando o defer vai fazer a função
	// ser executada por ultimo

	fmt.Println("fim")

	// o defer sempre executa depois de tudo, no final da
	// função

	// Ordem de execução (muito importante)

	/*
		Se você usar vários defer, eles funcionam como uma
		pilha (LIFO = último que entra, primeiro que sai):
	*/

	defer fmt.Println("1") // terceiro
	defer fmt.Println("2") // segundo
	defer fmt.Println("3") // vai ser primeiro
}
