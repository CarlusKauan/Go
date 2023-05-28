package main

import "fmt"

func main() {
	// endereço de memoria
	i := 1
	l := 2

	// Go não tem aritmética de ponteiros, refência de memoria
	var p *int = nil
	var q *int = nil
	// nil == null

	p = &i // pegando o endereço da variavel i e atribuindo ao p
	*p++   // incrementar o valor associado (pegando o valor)

	q = &l // pegando o endereço da variavel l e atribuindo ao q
	*q++   // incrementando o valor associado (pegando o valor)

	fmt.Println(p, *p, i, &i)
	fmt.Println(q, *q, l, &l)
}
