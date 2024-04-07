package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{ // Cria um map com a chave do tipo string e o valor é outro map com chave do tipo string e valor do tipo float64
		"G": {
			"Gabriela Silva": 15678.78,
			"Guga Pereira":   8567.89,
		},
		"J": {
			"José João": 11325.56,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P") // deleta todos os funcionários com a letra P

	for nome, salario := range funcsPorLetra { // Percorre por todo o map principal e retorn letra e funcs
		fmt.Println(nome, salario) // Printa o elemento atual (com dois pares de chave/valor)

		for nome, salario := range salario { // Percorre por todo o map principal e retorn letra e funcs
			fmt.Println(nome, salario) // Printa o elemento atual (com dois pares de chave/valor)
		}
	}
}
