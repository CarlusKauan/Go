package main

import "fmt"

// closure seria um fechamento entre as funções dadas...
func main() {
	x := 20
	fmt.Println(x)

	// criar uma variavel que representa a chamada da função closure
	imprimeX := closure()
	// Imprimir essa função
	imprimeX()
}

/* toda função é um closure, além do fato de sua função possuir um escopo, ela tem
a ideia de onde ela foi definida, a função traz seus valores ou parâmetros de onde foi
definida */

func closure() func() {
	// retornará uma função
	x := 10
	// atribuindo a uma função anonima
	var funcao = func() {
		fmt.Println(x)
	}
	// retorno de funcao
	return funcao
}
