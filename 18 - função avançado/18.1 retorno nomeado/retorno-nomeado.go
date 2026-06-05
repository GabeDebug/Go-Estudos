package main

import "fmt"

func calculoMatematico(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

// Função que recebe dois números inteiros (n1 e n2)
// e retorna dois valores nomeados:
// soma = n1 + n2
// subtracao = n1 - n2
// Os retornos são nomeados, então o "return" pode ser usado vazio.

func main() {
	soma, subtracao := calculoMatematico(10, 20)
	// aqui e como se fosse 10 + 20
	// e na subtração e 10 - 20
	fmt.Println(soma, subtracao)
}
