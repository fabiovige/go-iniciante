package main

import "fmt"

func main() {
	x := 5
	y := &x
	*y = 10
	fmt.Println("Valor de X = ", x, "Valor de Y = ", y)
	fmt.Println(&x, *y)
	imprimir(&x, y)
}

func imprimir(x *int, y *int) {
	fmt.Println("Função Imprimir")
	*x = 20
	fmt.Println("Ref de X = ", x, "Ref de Y = ", y)
	fmt.Println("Valor de X = ", *x, "Valor de Y = ", *y)
}
