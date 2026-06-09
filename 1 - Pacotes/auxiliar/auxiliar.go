package auxiliar

import (
	"fmt"
)

// Escrever registra uma mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}

// se uma função for com letra
// minuscula ela só vai ser vista no pacote que ela está
// com letra maiuscula ela vai ser exportada para outros pacotes
