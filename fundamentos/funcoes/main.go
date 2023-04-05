package main

/* Para usar funções dentro de outro arquivo é preciso criar modulos e criar funções públicas que são 
indicadas pela primeira letra maiúscula */
import (
	"funcoes/funcoes"
)

func main() {
	resultado := funcoes.Somar(3, 4)
	funcoes.Imprimir(resultado)
}
