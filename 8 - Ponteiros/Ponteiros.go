package main

import "fmt"

<<<<<<< HEAD
func main(){
=======
func main() {
>>>>>>> c1a3dff (Reconectando repositório após formatação)
	//Ponteiros
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	//Só vai mudar o valor da variavel 1
	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiros é referencia de memoria
<<<<<<< HEAD
	var variavel3 int 
=======
	var variavel3 int
>>>>>>> c1a3dff (Reconectando repositório após formatação)
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) // desreferenciação
	/*
<<<<<<< HEAD
	quando passo o * na frente eu estou dizendo 
	que vai nesse endereço de memoria
	*/
}
=======
		quando passo o * na frente eu estou dizendo
		que vai nesse endereço de memoria
	*/
}
>>>>>>> c1a3dff (Reconectando repositório após formatação)
