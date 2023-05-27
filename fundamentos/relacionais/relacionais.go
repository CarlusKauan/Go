package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	// times
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	// == comparação
	fmt.Println("Datas:", d1 == d2)
	// Equal usada para comparação
	fmt.Println("Datas:", d1.Equal(d2))

	// Structs
	type Pessoa struct {
		Nome string
	}

	type Automovel struct {
		nome   string
		placa  string
		modelo string
		kms    float64
	}

	a1 := Automovel{"Supra", "9AB23-0", "Toyota", 20345.23}

	fmt.Println("Modelo: "+a1.nome+"\nPlaca: "+a1.placa+"\nModelo: "+a1.modelo+"\nKms: ", a1.kms)

	p1 :=
		Pessoa{"João"}
	p2 :=
		Pessoa{"João"}

	fmt.Println("Pessoas:", p1 == p2)
}
