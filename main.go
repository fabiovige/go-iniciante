package main

import (
	"fmt"
	"time"

	"github.com/fabiovige/go-iniciante/model"
)

func main() {
	endereco := model.Endereco{
		Rua:    "Rua dos Bobos",
		Numero: 0,
		Cidade: "São Paulo",
	}

	pessoa := model.Pessoa{
		Nome:           "Fabio",
		Sobrenome:      "Vige",
		Email:          "fabiovige@gmail.com",
		Telefone:       "11999999999",
		DataNascimento: time.Date(1976, 03, 02, 0, 0, 0, 0, time.Local),
		Endereco:       endereco,
	}

	pessoa.CalculaIdade()
	fmt.Println("Idade ", pessoa.Idade)
}
