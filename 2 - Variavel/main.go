package main

import "fmt"

func main() {
	// tem dois tipo de variavel
	// deixando o var sem ser oculto
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	// deixando o tipo da variavel oculto
	variavel2 := "variavel2"
	fmt.Println(variavel2)

	// e tambem pode ser usada assim
	var (
		variavel3 string = "lalala"
		variavel4 string = "laksdjfk"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	// outro forma de fazer variável
	variavel5, variavel6 := "variavel5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "valor 1 da contante"
	fmt.Println(constante1)

	user := "relembrando os valores de variavel"
	fmt.Println(user)
}
