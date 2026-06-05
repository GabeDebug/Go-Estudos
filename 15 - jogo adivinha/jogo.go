package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da adivinhação")
	fmt.Println("Um número aleatoriio será sorteado. tente acerta.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual e seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("o seu chute tem que ser um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. o Número sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. o Número sorteado é Menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf("Parabéns voce acertou. o Numero era %d", x)
			return
		}

		chutes[i] = chuteInt
	}
	fmt.Printf("Infelizmente você não acertou um número %d", x)
}
