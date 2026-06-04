package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("resultado e 2")
	default:
		fmt.Println("Algo de errado")
	}
	// fmt.Println(isWeekend(time.Now()))
	//do(1)
}

func do(x int) {
	switch x {
	// posso tirar o X e colocar direto o valor
	case 1:
		fmt.Println(1)
		fallthrough // força a execução do próximo case
	case 2:
		fmt.Println(2)
		fallthrough
	default:
		fmt.Println("Outra coisa")
	}

}

func isWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}

}
