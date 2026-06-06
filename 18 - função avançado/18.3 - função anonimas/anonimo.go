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
	func(){}()

}
