package main

import "fmt"

func main() {
	fmt.Println("Inicio")

	defer fmt.Println("vai ser executado no final")
	/*
		Quando usamos o defer a função que estiver nela
		vai ser executada por ultimo
	*/

	fmt.Println("Vai ser executado em segundo")
}
