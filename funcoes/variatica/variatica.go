package main

import "fmt"

// Variatica você que determina a quantidade de valores/parâmetros manualmente !
func media(numeros ...float64) float64 {
	total := 0.0
	// for range para percorrer os parametros, acrescentando a soma !
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func somaValores(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func main() {
	fmt.Printf("Media do Aluno: %.2f \n", media(7.7, 8.1, 5.9, 9.9))
	fmt.Printf("A soma da pontuação Anual é: %d", somaValores(10, 10, 9))
}
