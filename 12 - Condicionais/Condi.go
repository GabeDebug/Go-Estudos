package main

import "fmt"

func main() {
	idade := 10

	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("você é menor de idade")
	}

	// usando else if
	notas := 4

	if notas >= 10 {
		fmt.Println("Exelente")
	} else if notas >= 7 {
		fmt.Println("bom")
	} else if notas == 5 {
		fmt.Println("recuperação")
	} else {
		fmt.Println("reprovado")
	}

	// declarando variavel dentro do if
	if idade2 := 20; idade2 >= 18 {
		// a variavel só vai existir dentro do bloco if
		fmt.Println("Maior de idade")
	}

	// operadores lógicos
	// AND &&
	temCarteira := true

	if idade >= 18 && temCarteira {
		fmt.Println("Pode dirigir")
	}

	// operador || OR
	if idade < 17 || !temCarteira {
		fmt.Println("Não pode Dirigir")
	}

	// operador ! NOT
	login := false

	if !login {
		fmt.Println("faça login")
	}
}
