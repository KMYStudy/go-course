package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("O Literal inteiro é", reflect.TypeOf(32000))

	// Sem sinal (só valores positivos)... uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))
	
	// Com sinal... int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de int é", reflect.TypeOf(i1))

	// Rune representa um mapeamento da tabela unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x)) // float32
	fmt.Println("O tipo de 49.99 é", reflect.TypeOf(49.99)) // float64

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá o meu nome é César"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// String com múltiplas linhas
	s2 := `Olá
	o meu nome 
	é César
	`
	fmt.Println("O tamanho da string é", len(s2))

	// Char ? Não existe !
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}