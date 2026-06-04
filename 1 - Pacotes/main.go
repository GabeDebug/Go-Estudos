package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("devbook@gmail.com")
	// vai me retorna que esse email é valido
	fmt.Println(error)
}