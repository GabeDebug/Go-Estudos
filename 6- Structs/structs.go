package main

import "fmt"

type usuario struct {
	Nome string
	Idade uint8
	endereco string
}

// structs e uma coleção de campo
func main(){
	// a forma mais recomendada para se trabalha com structs
	// primeira jeito
	usuario1 := usuario{Nome: "Gabriel", Idade: 30, endereco: "Rua dos bobos"}
	fmt.Println(usuario1)

	// segundo jeito
	var u usuario
	u.Nome = "Gabriel"
	u.Idade = 21
	fmt.Println(u)

	// terceiro jeito
	usuario3 := usuario{Nome: "davi"}
	fmt.Println(usuario3)

	//acessando e modificando campos
	// fmt.Println(p1.Nome)
	// p1.Idade = 21
}