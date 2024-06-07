package main

import "fmt"

func main() {

	helloMessage := "Hola"
	worldMessage := "Mundo"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(worldMessage, helloMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipos de Datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
