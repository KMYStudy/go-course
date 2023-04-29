package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 { // as chaves é obrigatório
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.2)
}
