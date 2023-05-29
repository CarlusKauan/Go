package main

import (
	"fmt"
)

// Encadeamento de condições
func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 7 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n <= 5 {
		return "D"
	} else {
		return "E"
	}
}

// Usando o switch
func notaParaConceitoSwitch(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 7 && n <= 8:
		return "B"
	case n >= 5 && n <= 6:
		return "C"
	case n >= 3 && n <= 4:
		return "D"
	case n >= 2:
		return "E"
	default:
		return "F"
	}
}

func main() {
	fmt.Println("Sua nota final:", notaParaConceitoSwitch(10))
	fmt.Println("Sua nota final:", notaParaConceitoSwitch(4.5))
	fmt.Println("Sua nota final:", notaParaConceitoSwitch(7.2))
	fmt.Println("Sua nota final:", notaParaConceitoSwitch(1.9))
}
