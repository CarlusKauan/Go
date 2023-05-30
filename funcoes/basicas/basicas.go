package main

import "fmt"

// func que não tem parametros e nem retorno
func f1() {
	fmt.Println("F1")
}

func f10() {
	fmt.Println("F10")
}

// Adicionando parametros
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f20(f21 string, f22 string) {
	fmt.Println("F20: %s, %s", f21, f22)
}

// Retorna uma string é obrigatorio !
func f3() string {
	return "F3"
}

func f30() string {
	return "F30"
}

// Declarando em sequencia ou serie
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f40(p41, p42 string) string {
	return fmt.Sprint(p41, p42)
}

// Quant. de retornos !
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func f50() (string, string) {
	return "Retorno 50", "Retorno 500"
}

// func principal
func main() {
	/* Função é um bloco de código que faz algo pré-definido ! */
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}
