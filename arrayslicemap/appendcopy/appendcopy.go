package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3} // Arrays geralmente tem tamanho fixo
	var slice1 []int // slice sem tamanho fixo e do tipo int (slices não tem tamanho fixo)

	//array1 = append(array1, 4, 5, 6) // o primeiro elemento precisa ser um array (não funciona desse jeito)
	slice1 = append(slice1, 4, 5, 6) // o append adiciona valores ao slice
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) // o slice será do tipo int e terá tamanho fixo de 2 elementos
	copy(slice2, slice1) // o copy copia tudo o que tem no slice1 para o slice2 (o copy não aumenta o tamanho do slice), ambos os valores precisam ser do mesmo tipo
	fmt.Println(slice2)
}
