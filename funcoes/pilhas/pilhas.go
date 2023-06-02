package main

// import que da condições da gente usar a função PrintStack()
import "runtime/debug"

func f3() {
	// Printa a sequência de funções ou as pilhas !
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

// porta de entrada a main
func main() {
	f1() // chamando pilhas de funções
	// stack estouro de pilhas
}
