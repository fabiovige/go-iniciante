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

	// exemplo de herança
	automovel := model.Automovel{
		Marca:  "Ford",
		Modelo: "Fiesta",
		Ano:    2019,
	}

	carro := model.Carro{
		Automovel: automovel,
		Portas:    4,
	}

	moto := model.Moto{
		Automovel:   automovel,
		Cilindradas: 125,
	}

	fmt.Println("Carro: ", carro)
	fmt.Println("Moto: ", moto.Ano)

	// exercicio 1
	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "Arroz")
	nomeDosItens = append(nomeDosItens, "Feijão")
	nomeDosItens = append(nomeDosItens, "Macarrão")

	compra, err := model.NovoPedido("Mercadinho", time.Now(), nomeDosItens)

	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Compra: ", compra)
	}
}
