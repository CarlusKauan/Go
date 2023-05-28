package main

import "fmt"

/* return nota >= 6 ? "Aprovado" : "Reprovado"
Não tem operadores ternários no golang por isso ussasse a
estrutura de código padrão de condições (if, else if)*/

// Não tem operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	} else if nota == 5 {
		return "Recuperação"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(10))
}
