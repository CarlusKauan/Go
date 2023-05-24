package main

import "fmt"

func main() {
	// importante a escolha ou declaração do tipo, fortemente tipada !
	var a int
	var b float64
	var c bool
	var d string
	// ponteiros de inteiro <nil> ou <null> por padrão !
	var e *int

	// %q equivalente as Strings que estão vazias !
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
	fmt.Println(" ")
	fmt.Printf("%v %q %v %v %v", e, d, c, b, a)
}
