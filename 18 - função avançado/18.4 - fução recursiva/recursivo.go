package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	//? estudando função recursivo
	fmt.Println("Estudos função recursivos")
	posicao := uint(10)
	fmt.Println(fibonacci(posicao))

}

/*
	1 1 2 3 5 8 13
	função recursiva é aquela que chama a si mesma, ou seja, ela se auto-invoca.
	Ela é composta por uma condição de parada e uma chamada recursiva.
	A condição de parada é essencial para evitar que a função entre em um loop infinito.

	Exemplo de função recursiva para calcular o fatorial de um número:
*/
