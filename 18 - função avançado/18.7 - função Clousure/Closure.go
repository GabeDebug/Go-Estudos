package main

import "fmt"

// Retorna uma closure que captura a variável 'texto' do seu escopo
func closure() func() {
	texto := "Dentro do closure" // a variavel que vai ser captura na main

	funcao := func() {
		fmt.Println(texto) // Acessa 'texto' capturado
	}

	return funcao
}

func alunos() func() {
	texto := "Aluno aprovado"

	funcao := func() {
		fmt.Println(texto) // teste usando função closure
	}

	return funcao
}

func main() {
	texto := "Dentro do Main"
	fmt.Println(texto)

	alunosAprovado := alunos()
	alunosAprovado()

	// Acessa 'texto' capturado
	funcao := closure()
	// Executa a closure que imprime o 'texto' capturad
	funcao() // Imprime oq está dentro do closure
}
