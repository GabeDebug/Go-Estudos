package main

import "fmt"

func main() {
	// Operadores Aritmeticos 
	// soma := 1 + 2
	//subtracao := 1 - 2
	// divisao := 10 / 2
	//multiplicacao := 10 * 5
	//restoDaDivizao := 10 % 2
	a := 10
	b := 3
	fmt.Println(a + b)

	// Operadores de Comparação

	fmt.Println(10 == 20) // false
	fmt.Println(10 != 5) // true
	fmt.Println(10 > 5) // true
	fmt.Println(10 < 5) // false
	fmt.Println(10 >= 10) // true
	fmt.Println(10 <= 8) // false
	/*
	| Operador | Significado    |
	| -------- | -------------- |
	| `==`     | Igual          |
	| `!=`     | Diferente      |
	| `>`      | Maior          |
	| `<`      | Menor          |
	| `>=`     | Maior ou igual |
	| `<=`     | Menor ou igual |
	*/

	// Operadores lógicos
	// Usando com valores booleanos
	idade := 20
	temCarteira := true
	
	fmt.Println(idade >= 18 && temCarteira)
	/*
	| Operador | Significado |
	| -------- | ----------- |
	| `&&`     | E (AND)     |
	| `\|\|`   | OU (OR)     |
	| `!`      | NÃO (NOT)   |
	*/
}