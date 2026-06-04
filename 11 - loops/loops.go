package main

import (
	"fmt"
)

func main() {
	var arr int
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(arr)

	// loop infinito

	// for {
	// 	fmt.Print("for")
	// }

	frutas := []string{"maçã", "banana", "uva"}

	for indice, fruta := range frutas {
		fmt.Println(indice, fruta)
	}

	numeros := []int{10,20,30}
	for indice, numeros := range numeros{
		fmt.Println(indice,numeros)
	}
}
