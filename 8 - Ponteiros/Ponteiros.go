package main

import "fmt"

func main(){
	//Ponteiros
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	//Só vai mudar o valor da variavel 1
	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiros é referencia de memoria
	var variavel3 int 
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) // desreferenciação
	/*
	quando passo o * na frente eu estou dizendo 
	que vai nesse endereço de memoria
	*/
}