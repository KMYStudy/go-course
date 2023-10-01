package main

import "fmt"

func main() {
	//var aprovados map[int]string // Esse map tem a chave do tipo int (cpf) e o valor do tipo string (nome)
	// mapas dever iniciailizados
	aprovados := make(map[int]string) // o make está inicializando o map

	aprovados[12345678978] = "Maria"
	aprovados[13456790754] = "Pedro"
	aprovados[11111111111] = "Carlos"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // Percorre todo o map aprovados e retorna o cpf e nome
		fmt.Printf("%s (CPF: %d)\n", nome, cpf) // Printa o nome e o cpf do aprovado
	}

	fmt.Println(aprovados[13456790754])
	delete(aprovados, 11111111111) // o delete exclui um valor do mapa, ele recebe dois parametros, um sendo o map e o outro a chave para ser excluída
	fmt.Println(aprovados[11111111111])
}
