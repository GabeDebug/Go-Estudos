package main

import (
	"fmt"
	"sync"
	"time"
)

func escrevernum(numeros int) {
	for i := 0; i < 2; i++ {
		fmt.Println(numeros)
		time.Sleep(time.Second)
	}
}

func main() {

	var WaitGroup sync.WaitGroup

	WaitGroup.Add(5)

	go func() {
		escrevernum(1)
		WaitGroup.Done()
	}()

	go func() {
		escrevernum(2)
		WaitGroup.Done()
	}()

	// todos vai ser executado em ordem aleatoria

	go func() {
		escrevernum(3)
		WaitGroup.Done()
	}()

	go func() {
		escrevernum(4)
		WaitGroup.Done()
	}()
	go func() {
		escrevernum(5)
		WaitGroup.Done()
	}()

	WaitGroup.Wait()

	fmt.Println("Fim do programa")

}
