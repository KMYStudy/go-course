package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{ // Cria o map e já defini os valores (chave: string, valor: float64)
		"José João":      11325.56,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0, // o último elemento tem que ter virgula
	}

	funcsESalarios["Rafael Filho"] = 1350.0 // Adiciona um novo elemento ao map
	delete(funcsESalarios, "Inexistente")   // Exclui um elemento que não existe no map (não acontecerá nada e não retornará nenhum erro)

	for nome, salario := range funcsESalarios { // percorre o map e retorna a key (nome) e o valor (salario) dos elementos
		fmt.Println(nome, salario) // Printa o elemento
	}
}
