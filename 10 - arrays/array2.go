package main

import "fmt"

var filmesNoDb = []string{
	"O Senhor dos Anéis: A Sociedade do Anel (2001)",
	"Star Wars – O Império Contra-Ataca (1980)",
	"O Poderoso Chefão (1972)",
	"Batman: O Cavaleiro das Trevas (2008)",
	"Um Sonho de Liberdade (1994)",
	"Tubarão (1975)",
	"Pulp Fiction: Tempo de Violência (1994)",
	"Vingadores: Guerra Infinita (2018)",
	"Os Caçadores da Arca Perdida (1981)",
	"Os Bons Companheiros (1990) ",
}

func main() {
	// slice vindo de uma resposta de uma api
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filmes := make([]string, 0, 10)
	// toda logica do db

	for _, id := range resultsFromApi {
		filme := filmesNoDb[id]
		filmes = append(filmes, filme)
	}

	fmt.Println(filmes)

}
