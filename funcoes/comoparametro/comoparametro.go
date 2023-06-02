package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) int {
	return a / b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	// Passando uma função como parametro !
	return funcao(p1, p2)
}

func exec2(funcao func(int, int) int, num1, num2 int) int {
	// Passando uma func como parametro
	return funcao(num1, num2)
}

func exec3(funcao func(int, int) int, num1, num2 int) int {
	return funcao(num1, num2)
}

// filter, map entre outras funções padronizadas

func main() {
	resultado := exec(mult, 3, 4)
	fmt.Println(resultado)
	// Recebendo uma função como parametros
	soma := exec2(sum, 10, 50)
	fmt.Println(soma)

	div := exec3(div, 50, 10)
	fmt.Println(div)
}
