package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	radius float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

func ExibeGeometria(g geometria) {
	fmt.Printf("Area do retangulo: %f\n", g.area())
}

func main() {
	fmt.Println("Inicializando...")
	r := retangulo{altura: 10, largura: 15}
	c := circulo{radius: 5}
	ExibeGeometria(r)
	ExibeGeometria(c)
}