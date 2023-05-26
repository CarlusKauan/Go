package main

import (
	"fmt"
	"math"
)

func main() {

	// Declaração e atribuição de variaveis
	a := 3
	b := 2

	// Operadores aritmeticos Básicos
	fmt.Println("OPERAÇÕES BASÍCAS")
	fmt.Println("Somar =>", a, "+", b, "=", a+b)
	fmt.Println("Subtrair =>", a, "-", b, "=", a-b)
	fmt.Println("Multiplicar =>", a, ".", b, "=", a*b)
	fmt.Println("Dividir =>", a, "/", b, "=", a/b)
	fmt.Println("Modulo =>", a, "%", b, "=", a%b)

	fmt.Print("\n")

	fmt.Println("OUTRAS OPERAÇÕES")

	// math libs
	c := 3.0
	d := 2.0

	fmt.Println("O maior é =>", math.Max(c, d))
	fmt.Println("O menor é =>", math.Min(c, d))
	// Exponenciação
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
