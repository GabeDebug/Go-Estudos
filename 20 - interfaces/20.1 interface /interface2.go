package main

import "fmt"

// ===== INTERFACE =====
// Uma interface é um "contrato" que diz: "qualquer tipo que queira ser Animal
// PRECISA ter esses dois métodos: Fazer() e Dormir()"
type Animal interface {
	Fazer()  // todo Animal precisa saber fazer um som
	Dormir() // todo Animal precisa saber dormir
}

// ===== STRUCT CACHORRO =====
// Cachorro é uma estrutura (tipo) que tem um campo Nome
type Cachorro struct {
	Nome string
}

// ===== IMPLEMENTANDO A INTERFACE PARA CACHORRO =====
// Este é um "método receptor" - a sintaxe é: func (receptor TipoDoReceptor) NomeDoMetodo()
// O "(c Cachorro)" significa: este método funciona COM um Cachorro
// "c" é uma variável que representa o Cachorro que está chamando este método

// Implementando o método Fazer() para Cachorro
// Quando você chama dog.Fazer(), este código aqui executa
func (c Cachorro) Fazer() {
	// "c.Nome" acessa o campo Nome do Cachorro
	// Isso imprime: "Rex faz: Au au!"
	fmt.Println(c.Nome + " faz: Au au!")
}

// Implementando o método Dormir() para Cachorro
// Quando você chama dog.Dormir(), este código executa
func (c Cachorro) Dormir() {
	// Imprime: "Rex está dormindo..."
	fmt.Println(c.Nome + " está dormindo...")
}

// ===== STRUCT GATO =====
// Gato também é uma estrutura com um campo Nome
type Gato struct {
	Nome string
}

// ===== IMPLEMENTANDO A INTERFACE PARA GATO =====

// Implementando o método Fazer() para Gato
// Mesma coisa que Cachorro, mas com comportamento diferente
func (g Gato) Fazer() {
	// "g.Nome" acessa o campo Nome do Gato
	fmt.Println(g.Nome + " faz: Miau!")
}

// Implementando o método Dormir() para Gato
func (g Gato) Dormir() {
	fmt.Println(g.Nome + " está dormindo...")
}

// ===== FUNÇÃO QUE RECEBE UMA INTERFACE =====
// Esta função é MUITO importante! Ela recebe um parâmetro do tipo "Animal"
// Isso significa: ela aceita QUALQUER coisa que implemente a interface Animal
// Pode ser Cachorro, Gato, ou qualquer outro tipo que tenha Fazer() e Dormir()
func FazerAnimalAtuar(animal Animal) {
	// Aqui chamamos os métodos da interface
	// Go não sabe se é Cachorro ou Gato, mas sabe que tem esses métodos!
	animal.Fazer()  // chama Fazer() do animal (seja qual for)
	animal.Dormir() // chama Dormir() do animal
}

// ===== MAIN - O PROGRAMA PRINCIPAL =====
func main() {
	// Criando um Cachorro chamado Rex
	// "{Nome: "Rex"}" é a forma de inicializar a struct com valores
	dog := Cachorro{Nome: "Rex"}

	// Criando um Gato chamado Whiskers
	cat := Gato{Nome: "Whiskers"}

	// ===== POLIMORFISMO EM AÇÃO! =====
	// Aqui está a magia das interfaces!
	// Estamos passando um Cachorro para uma função que espera um Animal
	// Go automaticamente reconhece que Cachorro implementa Animal
	// Porque Cachorro tem os métodos Fazer() e Dormir()
	FazerAnimalAtuar(dog)

	// Printando uma linha vazia para separar
	fmt.Println()

	// Agora estamos passando um Gato para a MESMA função
	// Mesmo código, comportamento diferente!
	// Isso é POLIMORFISMO - mesma função, resultados diferentes
	FazerAnimalAtuar(cat)
}
