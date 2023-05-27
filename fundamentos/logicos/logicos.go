package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv30 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv30, comprarSorvete
}

func main() {
	// Valores passados manualmente !
	tv50, tv32, sorvete := comprar(true, true)
	// Operadores unítarios !
	fmt.Printf("Tv50: %t, Tv32: %t, Saudável: %t", tv50, tv32, !sorvete)
}
