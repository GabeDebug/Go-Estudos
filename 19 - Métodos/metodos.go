package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados \n ", u.nome)
	// o %s serve para string
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
	//? aqui eu vou adicionar +1 na idade
}

func main() {
	//!usuario := usuario{"Usuario 1", 20}
	//!usuario.salvar()

	usuario2 := usuario{"Gabriel", 20}
	usuario2.salvar()

	//TODO aqui vai verificar se é maior de idade ou não
	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
