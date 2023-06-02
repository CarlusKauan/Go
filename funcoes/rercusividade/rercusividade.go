package main

import "fmt"

// Função que chama a sí propria !
func fatorial(n int) (int, error) {
	// normal em Golang o retorno do valor e o erro.
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número Invalido: %d", n)
	case n == 0:
		return 1, nil
	default:
		// momento da recursividade !
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	// Uma solução melhor seria...
}
