package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func Escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	Escrever("Estudando Golang", 1, 24, 1, 24, 12, 31, 25, 12, 3)
	totalDaSoma := soma(1, 2, 4, 5, 1, 5, 12, 3, 41)
	fmt.Println(totalDaSoma)
	escrever("Hello World", 1, 1, 1, 23, 5, 123, 1)
}
