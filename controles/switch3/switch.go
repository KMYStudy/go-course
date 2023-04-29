package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) { // verifica o tipo da variável i
	case int:
		return "Inteiro"
	case float32, float64:
		return "Número real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não conheço"
	}
}

func main() {
	fmt.Println(tipo(3.2))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
