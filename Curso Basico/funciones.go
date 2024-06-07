package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValuea(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Aloha Mundo")

	tripleArgument(1, 2, "Hola")

	value := returnValuea(2)
	fmt.Println("Valor: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Valor1 y Valor2:", value1, value2)
}
