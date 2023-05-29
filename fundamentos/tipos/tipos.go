package main

import (
	f "fmt"
	m "math"
	"reflect"
)

func main() {
	// números inteiros
	f.Println(1, 2, 1000)
	// Reflect exibe o tipo da variavel
	f.Println("Literal inteiro é:", reflect.TypeOf(32000))

	idade := 19.5
	f.Printf("%.2f", idade)
	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	var codigo byte = 111

	f.Println("O código final para rackear a russia: ", codigo)
	f.Println("O byte é: ", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := m.MaxInt64
	i2 := m.MaxInt32

	f.Println("O valor maximo de MaxInt32: ", i2)
	f.Println("O valor maximo do int é", i1)
	f.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i3 rune = 'a' // representa um mapeamento de Tabela Unicode (int32)
	f.Println("O rune é", reflect.TypeOf(i3))
	f.Println(i3)

	// números reais (float32, float64)
	var x = float32(49.99)
	f.Println("O tipo de x é", reflect.TypeOf(x))
	f.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	//boolean
	bo := true
	f.Println("O tipo de bo é", reflect.TypeOf(bo))
	f.Println("O inverso da variavel bo é", !bo)
	f.Println("O valor de bo é", bo)

	//Strings
	s1 := "Olá meu nome é Leo"
	my := "My name is Carlos"
	age := "I have nineteen years old"

	f.Println("-> ", my)
	f.Print("--" + age + "\n")

	f.Println(s1 + "!")
	f.Println("O tamanho da string é", len(s1))

	//String com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Carlos`
	f.Println(s2)
	f.Println("O tamanho da string é", len(s2))

	// char ???
	char := 'a'
	f.Println("O tipo de char é", reflect.TypeOf(char))
	f.Println(char)
}
