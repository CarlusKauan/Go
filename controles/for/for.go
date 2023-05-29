package main

import (
	"fmt"
	"time"
)

func main() {

	/* laço mais com uma quantidade determinada de repetições */
	i := 1
	for i <= 10 { // Não tem o while em go.
		fmt.Println(i)
		i++
	}
	// O for tem função de um while.

	for i := 0; i <= 20; i += 2 {
		// incremento i = i + 2, acrecentando dois em dois
		fmt.Printf("%d ", i)
	}

	/* poremos usar o for padrão, mas estranhamente o for pode ser usado
	como se fosse um while kk */

	fmt.Println("\n\nMisturando...")
	for i := 1; i <= 10; i++ {
		// Condição dentro de uma estrutura de for ou while querendo ou não
		if i%2 == 0 {
			fmt.Printf("%d é Par\n", i)
		} else {
			fmt.Printf("%d é Impar\n", i)
		}
	}

	for {
		// Laço infinito
		fmt.Println("Te amu vida ...")
		time.Sleep(time.Second)
	}
}
