package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch { // true
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	case t.Hour() > 19:
		fmt.Println("Boa noite")
	}
}
