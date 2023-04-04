package main

import "fmt"

func main() {
	// O comando Print não quebra a linha depois da execução
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// O comando Println quebra a linha depois da execução
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	//fmt.Println("O valor de x é:" + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é:", x)

	// Printf coloca ganchos no print
	fmt.Printf("O valor de x é %f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa!"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}