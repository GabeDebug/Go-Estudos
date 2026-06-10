package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) { // envia resposta ao cliente. / contém dados da requisição
	fmt.Fprintln(w, "Servidor Ok")
}

func main() {
	http.HandleFunc("/", home) // registra uma nova rota
	// "/" rota principal

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil) // inicia o servidor
}
