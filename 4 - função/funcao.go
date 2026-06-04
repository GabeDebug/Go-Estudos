package main

import "fmt"

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculoMatematicos(n1, n2 int8) (int8, int8){
		soma := n1 + n2
		subtracao := n1 - n2
		return soma, subtracao
}

func saudacao(nomes string) string {
	return "Olá, " + nomes + "!"
}

func main() {
	somar := soma(10,20)
	fmt.Println(somar)

	var f = func(texto string) {
		fmt.Println(texto)
	}
	f("texto da função F")

	// resultado vai ser 10 + 15 vai ser 25
	// e 10 - 15 vai ser -5
	resultadoSoma, resultadoSubtracao := calculoMatematicos(10,15)
	// passando o _ eu sou vou pega o resultado da Primeira conta
	fmt.Println(resultadoSoma,resultadoSubtracao)

	resultado := saudacao("Gabriel")
	fmt.Println(resultado)
}