package main

import "fmt"

// val soma recebe resultado de função anonima !
var soma = func(a, b int) int {
	return a + b
}

func main() {
	// Programação funcional !
	fmt.Println(soma(10, 50))
	// forma de declaração/atribuição de val que será armazenado pelo resultado da função.
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println("SUBTRAÇÃO: ", sub(10, 50))

	mult := func(a, b int) int {
		return a * b
	}
	fmt.Println("MULTIPLICAÇÃO: ", mult(10, 50))

	div := func(a, b int) int {
		return a / b
	}
	fmt.Println("DIVISÃO: ", div(50, 10))
}
