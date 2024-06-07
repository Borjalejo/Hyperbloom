package main

import (
	pk "Curso_Golang/src/mypackage" //Alias pk
	"fmt"
)

func main() {
	var micarro pk.CarPublic
	micarro.Marca = "Ferrari"
	micarro.Anio = 2020
	fmt.Println(micarro)

	pk.MostMensaje()
}
