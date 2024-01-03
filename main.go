package main

import (
	"errors"
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

	// implementação da interface error
	ExibeErro(errors.New("novo erro de warning"))

	p := ProblemaDeNetwork{
		rede: false,
		hardware: true,
	}
	ExibeErro(p)

	// interface vazia

	var lista []interface{}
	lista = append(lista, "Fabio")
	lista = append(lista, 1)
	lista = append(lista, true)

	for _, item := range lista {
		fmt.Println(item)

		v, ok := item.(string)
		if ok {
			fmt.Println("É uma string: ", v)
		} else {
			fmt.Println("Não é uma string")
		}

	}

}

type ProblemaDeNetwork struct {
	rede bool
	hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	problemas := ""
	if p.rede {
		problemas += "rede "
	}
	if p.hardware {
		problemas += "hardware "
	}
	return problemas
}

func ExibeErro(err error) {
	fmt.Println("Erro: ", err.Error())
}
