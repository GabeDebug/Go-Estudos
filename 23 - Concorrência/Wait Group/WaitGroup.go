package main

import (
	"fmt"
	"sync"
	"time"
)

// Sincronizando a goroutines

func main() {
	// Cria uma WaitGroup - ela funciona como um contador de tarefas
	var WaitGroup sync.WaitGroup

	// Adicionamos 2 ao contador porque vamos lançar 3 goroutines
	WaitGroup.Add(3)

	go func() {
		escrever("Olá mundo")
		WaitGroup.Done() // Informa que esta goroutine terminou
		/*
			Quando passamos o Done ele tira -1 do contator
		*/
	}()

	go func() {
		escrever("programando em Golang")
		WaitGroup.Done() // Informa que esta goroutine terminou
		/*
			Quando passamos o Done ele tira -1 do contator
		*/
	}()

	go func() {
		escrever("Fim do programa")
		defer WaitGroup.Done() // posso coloca o defer no final para garantir que seja chamado mesmo se ouver error
	}()

	WaitGroup.Wait()

	fmt.Println("Todos os goroutines finalizaram!")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
