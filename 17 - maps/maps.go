package main

import (
	"fmt"
)

func main() {
	// estudos sobre map
	fmt.Println("Maps")

	// dentro do cochetes e o tipos da chave
	// e fora e tipo do valores
	usuario := map[string]string{
		"nome":      "gabriel",
		"sobrenome": "Almeida",
	}

	/*
		estou passando um map com valor de string
		e passando o segundo map com mesmo valor
	*/
	usuario2 := map[string]map[string]string{
		"curso": {
			"Golang": "Golang básico",
		},
		"Faculdade": {
			"Faculdade": "Estácio",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "curso")
	// deletando o valor dentro da chaves
	fmt.Println(usuario2)

	// para adicionar
	usuario2["Signo"] = map[string]string{
		"Nome": "áries",
	}

	fmt.Println(usuario2)

	fmt.Println(usuario)
	// pra acessa eu passo [""] ai consigo acessar
	/*
		o map em Go serve para fazer busca rápidas por chave
		quando você precisa associar um valor a um identificador
	*/

	nomes := map[string]int{
		"Gabriel": 1,
		"joão":    2,
		"Maria":   3,
	}

	delete(nomes, "Gabriel")
	// aqui vou deleta o valor que está dentro do map

	fmt.Println(nomes)

	nomes2 := map[int]string{
		1: "Gabriel",
		2: "joão",
		3: "Maria",
	}

	fmt.Println(nomes2)

}
