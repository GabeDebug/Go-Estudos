package main

import "fmt"

func main() {
	// estudos sobre map
	idades := map[string]int{
		"joão":  25,
		"maria": 26,
		"jose":  27,
	}

	// idades["ana"] = 25
	// idades["João"] = 30

	fmt.Println(idades["joão"])
	/*
		o map em Go serve para fazer busca rápidas por chave
		quando você precisa associar um valor a um identificador
	*/
}
