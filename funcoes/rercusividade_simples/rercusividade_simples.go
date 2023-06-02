package main

import "fmt"

// uint valor inteiro sem decimal
func fatorial(n uint) uint {
	// normal em Golang o retorno do valor e o erro !
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado1 := fatorial(5)
	fmt.Println(resultado1)

	// daria um erro pois não aceita negativos (n uint) uint
	// resultado2 := fatorial(-4)
	// fmt.Println(resultado2)
}
