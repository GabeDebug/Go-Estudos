package main

import "fmt"

func main() {
	//? estudando sobre função anonima
	fmt.Println("Função anonima")

	func(texto string) {
		fmt.Println(texto)
	}("passando parametro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Passando -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)

	retorno2 := func(texto string) string {
		return fmt.Sprintf("Passando o -> %s", texto)
	}("Passando parametro 2")

	fmt.Println(retorno2)

	//! sintaxe da função anonima
	func() {}()

	//? como criar uma função anonima usando soma
	numeros2 := func(a, b int) int {
		return a + b
	}
	resultado := numeros2(1, 1)
	fmt.Println(resultado)
}
