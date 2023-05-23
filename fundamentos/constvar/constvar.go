package main

import (
	f "fmt"
	// Referência as libs
	m "math"
)

// import mais de uma lib

func main() {

	// tipo (float64) inferido pelo computador
	const PI float64 = 3.1415
	// declarando um var
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar um var e inicializar usando (:=)
	area := PI * m.Pow(raio, 2)
	f.Printf("Área de circusferência: %.2f", area)

	const (
		a = 1
		b = 2999
	)

	/*
		Exemplo de declaração de variavéis
		e constantes de modo que você pode declarar varios
		em sequência !
	*/

	var (
		c = 12
		d = 10
	)
	// tipo de declaração seguindo uma ordem de valores !
	f.Println(a, b, c, d)

	// Outro modo de declaração e atribuição em sequência.
	g, h, i := 2, false, "epa"

	f.Println(g, h, i)

}
