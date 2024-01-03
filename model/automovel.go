package model

type Automovel struct {
	Marca  string
	Modelo string
	Ano    int
}

type Carro struct {
	Automovel
	Portas int
}

type Moto struct {
	Automovel
	Cilindradas int
}
