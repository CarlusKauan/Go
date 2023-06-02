package main

import "fmt"

// Parâmetros variaticos estilo string
func imprimir(aprovados ...string) {
	// trata-se parâmetros como uma array...
	for i, aprovados := range aprovados {
		// percorrer o array aprovados
		fmt.Printf("%d) %s\n", i+1, aprovados)
	}
}

func main() {
	// Array que o compilador vai contar os parâmetros, mas funciona com slice
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme", "Carlos"}
	// imprimir a Array
	imprimir(aprovados...)
}
