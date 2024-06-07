package main

import "fmt"

func main() {
	m := make(map[string]int)

	//Agregar valor de una persona
	m["Jose"] = 10
	m["Pepito"] = 20
	fmt.Println(m)

	//Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
