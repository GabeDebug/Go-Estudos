package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (Dog) Sound() string {
	return "Au AU"
}

func somDoAnimal(a Animal) {
	fmt.Println(a.Sound())
}

func main() {
	// Trabalhando com Type Assertions

	dog := Dog{}
	somDoAnimal(dog)
}
