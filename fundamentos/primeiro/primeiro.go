/*
Programa executáveis iniciam pelo pacote main,
que será chamado posteriomente no decorrer do código
*/
package main

/*
Os codigos em Go são organizados em pacotes
e para usa-los é necessario declarar um ou varios imports
como o 'import fmt'
*/
import "fmt"

// A porta de entrada de um programa Go é a função main
func main() {

	// fmt.Println é ultilizado para imprimir algo na tela, e pular uma linha.
	message := "Hello World !"
	fmt.Printf("%v", message)
}

/*
	Sobre comentários...

	1) Priorize código legível e faça comentários que agrega valor.
	2) Evite comentários óbvios.
	3) Durante o curso abuse dos comentários.

*/
