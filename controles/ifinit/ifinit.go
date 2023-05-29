package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	// rand - random
	s := rand.NewSource(time.Now().UnixNano()) // data atual, nano segundo
	r := rand.New(s)                           // criar número aleatorio
	return r.Intn(10)                          // retornar a fonte, e gera um número até 10
}

func main() {
	if i := numeroAleatorio(); i > 5 { //tb suportado no
		fmt.Println("Ganhou, número:", i)
	} else {
		fmt.Println("Perdeu, número:", i)
	}
}

// Veremos foreach no capitulo de array
