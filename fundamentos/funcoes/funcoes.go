package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func subtrair(a int, b int) int {
	return a - b
}

func multiplicar(a int, b int) int {
	return a * b
}

func dividir(a int, b int) int {
	return a / b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	somar := somar(3, 4)
	subtrair := subtrair(10, 5)
	multiplicar := multiplicar(3, 3)
	dividir := dividir(50, 2)
	imprimir(somar)
	imprimir(subtrair)
	imprimir(multiplicar)
	imprimir(dividir)
}
