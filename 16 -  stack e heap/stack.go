package main

import "fmt"

func main() {
	// trabalhando com ponteiro
	// Stack e Heap
	x := 100
	p := &x            // ele guarda o valor de x
	fmt.Println(x, *p) // aqui ele vai até aonde tá guardado e me retorna

	a := 10
	alterar(&a)
	fmt.Println(a)
}

// para modificar valores fora da função
func alterar(x *int) {
	*x = 99
}
