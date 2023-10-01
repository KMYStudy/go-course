package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for i, numero := range numeros { // Ele vai percorrer pelo array numeros
		fmt.Printf("%d) %d\n", i + 1, numero) // Printara todos até chegar o final 
	}

	// _ é usado para ignorar uma variável
	for _, num := range numeros {
		fmt.Println(num)
	}
}
