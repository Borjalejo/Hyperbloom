package main

import "fmt"

func main() {

	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	//Division
	result = y / x
	fmt.Println("Division: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Retos
	// Rectangulo, Trapecio y un circulo

	//Area Rectangulo
	var largo float64 = 7.5
	var ancho float64 = 15
	areaRectangulo := largo * ancho
	fmt.Println("El area de un rectangulo es de: ", areaRectangulo)

	//Area Trapecio
	var bMayor float64 = 9.5
	var bMenor float64 = 3.5
	var Altura float64 = 4
	areaTrapecio := ((bMayor + bMenor) / 2) * Altura
	fmt.Println("El area de un trapecio es de: ", areaTrapecio)

	//Area Circulo
	var radio float64 = 8
	areaCirculo := pi2 * (radio * radio)
	fmt.Println("El area de un Circulo es de: ", areaCirculo)
}
