package main

import "fmt"

type pessoa struct{
	Nome string
	Sobrenome string
	Idade uint8
	Altura uint8
}

type estudante struct{
	// fazendo assim eu vou está pegando todos os campos
	// que está dentro de pessoa e passando pra estudante
	pessoa
	Curso string
	Faculdade string
}

func main() {
	pessoa1 := pessoa{"Gabriel", "Almeida", 20, 190}
	fmt.Println(pessoa1)
	// para eu conseguir criar um estudante preciso criar uma 
	// pessoa antes pra conseguir fazer isso
	
	estudante1 := estudante{pessoa1, "Engenharia de software", "Usp"}
	fmt.Println(estudante1)
}