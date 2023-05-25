package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	z := 10
	// Você precisa fazer a conversão de um tipo que seja adequado para determinadas conversões !
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)
	fmt.Println((int(x) * y) + z)

	nota := 6.9
	// Conversão feito que arredonda o number
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97)) // Tabela ASC

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Test two" + strconv.Itoa(123))
	fmt.Println("Teste three" + strconv.Itoa(123))
	fmt.Println("Teste four" + strconv.Itoa(100))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

	// Sistema de Notas
	nota_um := 5
	nota_dois := 7
	nota_tres := 10
	totalnotas := float64(nota_um) + float64(nota_dois) + float64(nota_tres)
	media := (totalnotas) / 3
	fmt.Printf("%.2f", media)
}
