package main

import "runtime/debug"

/*

	Em resumo, dependendo do contexto, "pilha" em Go pode se referir a uma estrutura de dados implementada usando fatias ou 
	à pilha interna usada para rastrear chamadas de função e a execução de goroutines.

	Uma nova função só é chamada quando a função anterior terminou a sua execução

*/

func f3() {
	debug.PrintStack() // Retorna a pilha de funções que foram chamadas no momento que foi executado
}

func f2() {
	f3() // Antes da f2 acabar, ela chama a f3
}

func f1() {
	f2() // Antes da f1 acabar, ela chama a f2
}

func main() {
	f1() // Executa a f1
}
