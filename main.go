package main

import "fmt"

type endereco struct {
	rua    string
	numero int
	cidade string
}

func main() {
	fmt.Println("Hello, World!")

	endereco := endereco{
		"Rua dos Bobos",
		0,
		"São Paulo",
	}

	fmt.Println(endereco)
	endereco.numero = 449
	fmt.Println(endereco)
}
