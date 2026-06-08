package main

import "fmt"

var n int
var p string

func init() {
	fmt.Println("Executando a função init")
	n = 10
	p = "Gabriel"
	// a função init vai ser executada primeiro que a função main
}

func main() {
	//? Estudando sobre função init
	fmt.Println("Executando a função main")
	fmt.Println(n)
	fmt.Println(p)
}
