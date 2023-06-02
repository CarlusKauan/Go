package main

// Para executar tudo usasse, go run .
import "fmt"

// Init será executado primeiro que main !
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
