package main

import "fmt"

func main() {

	fmt.Printf("Alguns Comandos \n")
	fmt.Printf("Comandos em %s!!!!\n", "Go")
	fmt.Printf("Comandos em %s \n", "Ruby")
}

// Utilização de comandos de go
/*
	$ go
	$ go help get
	$ go version
	$ godoc -http=:6060
	$ go env
	$ go doc cmd/vet
	$ go vet comandos.go
	$ go build comandos.go
	$ ./comandos

	# Testar arquivo.go
	$ go run comandos.go

	# Mac/Linux
	$ ls ~/go/src/github.com
	# Windows
	$ dir ~/go/src/github.com

	# Instalar pacote mysql
	$ go get -u github.com/go-sql-driver/mysql
*/
