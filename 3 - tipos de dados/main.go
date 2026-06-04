package main

import (
	"errors"
	"fmt"
)

func main() {
	// todos tipos básicos
	//int8, int16,int32,int64 todos os tipo de int
	var numero int64 = 23871
	fmt.Println(numero)

	//var numerop2 uint32 = -1000
	//fmt.Println(numerop2) //unsygned int e um int sem sinal

	// alias
	//int32 = Rune
	// posso usar o rune no lugar do int32 vai ser o mesmo valor
	var numero3 rune = 14124
	fmt.Println(numero3)

	//Byte = Uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// numeros reais
	var numeroReal float32 = 12.353
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 12.23487353
	fmt.Println(numeroReal2)

	numeroReal4 := 32.384
	fmt.Println(numeroReal4)

	// FIM NÚMEROS REAIS
	var str string = "kasdjklfa"
	fmt.Println(str)

	// o char não tem no golang
	char := 'b'
	fmt.Println(char)

	//FIM STRING
	texto := 5
	fmt.Println(texto)

	// boolean / erro
	var boolean1 bool = true
	fmt.Println(boolean1)

	// erro e o nome var
	// error com R e o tipo
	//e errors e o nome do pacote
	var erro error = errors.New("erro interno")
	fmt.Println(erro)
	// vai me retorna o valor nill 
}