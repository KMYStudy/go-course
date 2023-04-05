package main

import "fmt"

func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado" // O Go não suporta operador ternário
	
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}