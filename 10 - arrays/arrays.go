package main

import (
	"fmt"
)

func main() {
	arr := [3]int{10, 20, 30}
	arr2 := [10]int{1: 100, 9: 200}
	// eu posso fazer todo tipo de array

	nomes := [5]string{1: "gabriel", 0: "davi"}
	fmt.Println(nomes)

	fmt.Println(arr2)

	// usando array com indices
	arr[2] = 100
	fmt.Println(arr)

	// slices

	a := [5]int{1, 2, 3, 4, 5}
	slice := a[1:4]
	a[1] = 999
	// mudando o array que tbm vai mudar o slice
	slice[0] = 123
	// mudando o slice que tbm vai mudar no array

	fmt.Println(a)

	// Declarando um slice sozinho

	slice2 := []int{1, 2, 3}
	// e a mesma coisa do array mais eu não preciso
	// passa o tamanho dele
	fmt.Println(slice2)
	// o limite inferior tem por padrão 0
	// e o limite superior tem por padrão o len do array

	fmt.Println(slice2, len(slice2), cap(slice2))
	// esse slice tem capacidade de 3 elemento
	// e está capacitando 3 elemento
}
