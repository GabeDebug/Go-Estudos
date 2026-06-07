package main

import "fmt"

func funcao1() {
	fmt.Println("executando na tela 1 ")
}

func funcao2() {
	fmt.Println("executando na tela 2 ")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. resutlado será retornado")
	// vai ser executado antes do return independente quem seja!

	fmt.Println("aprovado ou não?")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
		// aqui vai ser aprovado
	}

	return false
	// aqui vai ser reprovado
}

func calculor(a, b int) int {
	defer fmt.Println("Está sendo calculado.")
	// vai ser executado antes do return

	soma := a + b
	return soma
}

func main() {

	resultadoSoma := calculor(1, 5)
	fmt.Println(resultadoSoma)

	resultado := alunoAprovado(7, 8)
	fmt.Println(resultado)

	// DEFER => ADIAR
	// defer funcao1()
	// funcao2()
}
