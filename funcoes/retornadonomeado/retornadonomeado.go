package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) { // O retorno quando é nomeado, ele se comporta como se fosse variáveis que são retornadas
	segundo = p2
	primeiro = p1
	return // Retorno limpo, dessa forma não precisa retornar as variáveis
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
