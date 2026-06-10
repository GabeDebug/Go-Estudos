package main

import (
	"fmt"
	"strings"
)

func main() {

	texto := "Testando a lib string"

	// Converter para maiúsculo
	fmt.Println(strings.ToUpper(texto))

	// para conta caracteres
	fmt.Println(strings.Count(texto, "T"))

	// Dividir por espaço
	palavras := strings.Split(texto, " ")
	fmt.Println(palavras)

	//Substituir
	novo := strings.Replace(texto, "lib", "Go", 1)
	fmt.Println(novo)

	//Clone da string
	new := strings.Clone(texto)
	fmt.Println(new)

	//Ver se contem a string
	novo2 := strings.Contains(texto, "Testando")
	fmt.Println(novo2)

	//Deixa a string em letra minuscula
	novo3 := strings.ToLower(texto)
	fmt.Println(novo3)

	//Adicionar o valor dentro string
	novo4 := []string{"1", "2", "4"}
	resultado := strings.Join(novo4, ",")

	fmt.Println(resultado)
}
