package main

import "fmt"

// retorno nomeado no (retorno int, retorno int)
func trocar(p1, p2 int) (segundo int, primeiro int) { // (segundo, primeiro int)
	segundo = p2
	primeiro = p1
	return // retorno limpo * return segundo, primeiro

}

func trocarNome(name, name2 string) (segundo string, primeiro string) { // (segundo, primeiro string)
	segundo = name2
	primeiro = name
	return
}

func main() {
	fmt.Println(trocar(10, 20))
	fmt.Println(trocarNome("Carlos", "Kauan"))

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
