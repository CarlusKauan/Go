package main

import "fmt"

func main() {
	// estrutura homogenea (mesmo tipo) e estatica (fixo)
	var notas [3]float64
	// estrutura index
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1 // atribuição de notas
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)

	var carros [3]string
	carros[0], carros[1], carros[2] = "Honda", "Toyota", "Civic"
	fmt.Println(len(carros), carros)

	var names [2]string
	names[0], names[1] = "Carl", "Emily"
	fmt.Println(len(names), names)
}
