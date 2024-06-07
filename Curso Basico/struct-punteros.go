package main

import "fmt"

type pc struct {
	ram   int
	disco int
	marca string
}

func (compu pc) ping() {
	fmt.Println(compu.marca, "Pong")

}

func (compu *pc) dobram() {
	compu.ram = compu.ram * 2

}

func main() {

	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	compu := pc{ram: 16, disco: 200, marca: "msi"}
	fmt.Println(compu)

	compu.ping()

	fmt.Println(compu)
	compu.dobram()

	fmt.Println(compu)
	compu.dobram()

	fmt.Println(compu)
}
