package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		//Se existe um erro capturado (r não é vazio), entra no bloco
		fmt.Println("Execução recuperada")
	}
	//Recover = Captura o erro gerado por panic()
}

func mediaAluno(n1, n2 float64) bool {
	//defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA FOI EXATAMENTE 6")
	/*
		vai me retorna um error e vai interromper meu codigo
		e posso chama o recover para fazer tratamento de error
	*/
}

func main() {

	fmt.Println(mediaAluno(6, 6))
	fmt.Println("Pós execução!")
	// aqui vai ser a pós execução quando eu chama ela
}

/*
1. main() chama mediaAluno(6, 6)
2. defer recuperarExecucao() é registrado
3. media = 6
4. Não entra em nenhuma condição (não > 6 e não < 6)
5. panic() é chamado! ❌
6. defer executa: recuperarExecucao() captura o erro ✅
7. "Pós execução!" NÃO é imprimido (programa interrompe)
*/
