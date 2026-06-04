package main

import "fmt"

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
}
