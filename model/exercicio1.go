package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado    string
	DataCompra time.Time
	Itens      []Item
}

type Item struct {
	Nome string
}

func NovoPedido(mercado string, dataCompra time.Time, nomeDosItens []string) (*Compra, error) {

	if mercado == "" {
		return nil, errors.New("mercado não pode ser vazio")
	}

	var itens []Item

	for _, nome := range nomeDosItens {
		itens = append(itens, Item{Nome: nome})
	}

	return &Compra{
		Mercado:    mercado,
		DataCompra: dataCompra,
		Itens:      itens,
	}, nil
}
