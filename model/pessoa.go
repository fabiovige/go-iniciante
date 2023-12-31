package model

import (
	"time"
)

type Pessoa struct {
	Nome           string
	Sobrenome      string
	Email          string
	Telefone       string
	DataNascimento time.Time
	Idade          int
	Endereco       Endereco
}

func (pessoa *Pessoa) CalculaIdade() {
	pessoa.Idade = time.Now().Year() - pessoa.DataNascimento.Year()
}

func CalculaIdade(pessoa Pessoa) int {
	return time.Now().Year() - pessoa.DataNascimento.Year()
}
