package main

import "fmt"

// Copia do valor
func invertido(numero int) int {
	return numero * -1
	// criação de função invertida
	// estamos passando uma copia pra esse função
	/*
		 'numero' é uma CÓPIA do valor passado
		Qualquer mudança aqui NÃO afeta a variável original
		Exemplo: se chamar invertido(5), número recebe uma cópia de 5
		se eu quiser fazer uma copia eu uso
	*/
}

func valorNaoAlterado(numeros int) int {
	return numeros * -1
}

// alterando o valor da variavel com ponteiro
func invertePonteiro(numero *int) {
	*numero = *numero * -1
	// quando eu faço isso que dizer que eu estou indo pega o valor
	// e alterando seu sinal (colocando menos)
}

func inverterSinalComPonteiro(numeros *int) {
	*numeros = *numeros * -1
	/*
		passando um ponteiro está passando uma referencia
		para essa função é qualquer alteração feita na variavel
		vai afeta o valor /
		se eu quiser realmente alterar o valor da variavel no meu codigo
		eu uso ponteiro
	*/
}

func main() {

	newNumber := 30
	numeroN := valorNaoAlterado(newNumber)
	fmt.Println(numeroN)

	//? estudos sobre Função de ponteiro
	numero := 20
	numeroInvertido := invertido(numero)
	// copia da variavel e manda pra função
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	// passei o endereço de memoria
	fmt.Println(novoNumero)
}

// '*int' significa: "recebe um PONTEIRO para um inteiro"
// Um ponteiro é um endereço de memória (referência)

// '*numeros' desreferencia o ponteiro
// Ou seja, acessa o VALOR que está naquele endereço de memória

// '*numeros * -1' pega o valor apontado, multiplica por -1
// e armazena DE VOLTA no mesmo endereço de memória
// Isso ALTERA a variável original!

/* ===== EXEMPLO 1: Passando uma CÓPIA =====
Antes: numero1 = 10
Resultado: -10
Depois: numero1 = 10

Por que não mudou?
→ invertido() recebeu uma CÓPIA de 10
→ A cópia foi multiplicada por -1
→ A variável original numero1 continua 10

===== EXEMPLO 2: Passando um PONTEIRO =====
Antes: numero2 = 10
Depois: numero2 = -10

Por que mudou?
→ inverterSinalComPonteiro() recebeu um PONTEIRO (endereço)
→ Com *numeros, ele acessou o valor no endereço
→ Alterou o valor naquele endereço
→ A variável original numero2 foi modificada!

===== RESUMO DA DIFERENÇA =====
CÓPIA (valor):     invertido(numero) → não altera original
REFERÊNCIA (ponteiro): inverterSinalComPonteiro(&numero) → altera original*/
