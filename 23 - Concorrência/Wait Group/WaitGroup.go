package main

import (
	"fmt"
	"sync"
	"time"
)

// Sincronizando a goroutines

func main() {
	var WaitGroup sync.WaitGroup

	WaitGroup.Add(3)

	go func() {
		escrever("Olá mundo")
		WaitGroup.Done()
		/*
			Quando passamos o Done ele tira -1 do contator
		*/
	}()

	go func() {
		escrever("programando em Golang")
		WaitGroup.Done()
		/*
			Quando passamos o Done ele tira -1 do contator
		*/
	}()

	go func() {
		escrever("Fim do programa")
		defer WaitGroup.Done() // posso coloca o defer no final para garantir que seja chamado mesmo se ouver error
	}()

	WaitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
