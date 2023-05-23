package main

import "fmt"

func main() {

	// Print normal
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// Print pulando uma linha
	fmt.Println(" Nova")
	fmt.Println("linha.")

	// Declaração simplificada :=
	x := 3.141516

	// Retorna o valor, qualquer valor em String
	xs := fmt.Sprint(x)

	// Transformado em String e reescrito no fmt
	fmt.Println("(String) O valor de x é " + xs)
	fmt.Println("(Number) O valor de x é", x)

	// Outra maneira de formatar personalizamente
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	// Printf outra maneira de imprimir com os simbolos correspodentes !
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// Printf usando %v que pode interpretar qualquer tipo de variaveis
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
