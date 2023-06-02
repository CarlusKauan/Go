package main

import "fmt"

// atrasa um pouco a função - defer
func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor Alto")
		return numero
	}
	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
