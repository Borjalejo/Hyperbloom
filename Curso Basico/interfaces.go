package main

import "fmt"

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

type fig2D interface {
	area() float64
}

func calcu(f fig2D) {
	fmt.Println("Area: ", f.area())
}

func main() {
	micuadrado := cuadrado{base: 2}
	mirectangulo := rectangulo{base: 2, altura: 4}

	calcu(micuadrado)
	calcu(mirectangulo)

	//Lista de interfaces
	miinter := []interface{}{"Hola", 12, 25.4}
	fmt.Println(miinter...)

}
