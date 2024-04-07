package main

import "fmt"

func f1() { // Função comum que não passa parametros e nem retorna nada
	fmt.Println("F1") // Apenas printa a mensagem
}

func f2(p1 string, p2 string) { // Funções que passam parametros deve ter os tipos de parametros definidos e não retorna nada
	fmt.Printf("F2: %s %s\n", p1, p2) // Apenas printa a memsagem com o Printf
}

func f3() string { // Funções que retornam algo deve ter o tipo de retorno definido (no caso, o retorno é do tipo string)
	return "F3" // Retorna um string
}

func f4(p1, p2 string) string { // Parametros do mesmo tipo pode ser definido apenas uma vez
	return fmt.Sprintf("F4: %s %s", p1, p2) // Apenas retorna a mensagem com um Sprintf
}

func f5() (string, string) { // Quase a função retorna multiplos valores, se usam parenteses com os tipos dos retornos dentro
	return "Retorno 1", "Retorno 2" // Retorna as duas strings
}

func main() {
	f1()
	f2("Param1", "Param 2")

	r3, r4 := f3(), f4("Param 1", "Param 2") // Em Go podem se chamar duas funções ao mesmo tempo nas variáveis
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5() // Quando a função retorna dois ou mais valores, você deve tratar os retornos (mas, se não usar variáveis, o tratamento pode ser ignorado), com _ você ignora um retorno
	fmt.Println("F5:", r51, r52)
}