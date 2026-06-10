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
}
