package main

import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	arr2 := [10]int{1: 100, 9: 200}

	fmt.Println(arr2)

	// indices
	arr[2] = 100
	fmt.Println(arr)
}
