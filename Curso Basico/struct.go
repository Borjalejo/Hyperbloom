package main

import "fmt"

type car struct {
	marca string
	anio  int
}

func main() {
	carro := car{marca: "Ford", anio: 2020}
	fmt.Println(carro)

	//Otra manera
	var otroCarro car
	otroCarro.marca = "Ferrari"
	fmt.Println(otroCarro)
}
