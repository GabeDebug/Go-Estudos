package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	totalDaSoma := soma(1, 2, 4, 5, 1, 5, 12, 3, 41)
	fmt.Println(totalDaSoma)
}
