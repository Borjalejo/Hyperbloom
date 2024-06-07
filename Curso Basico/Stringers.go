package main

import "fmt"

type pc struct {
	ram   int
	marca string
	disco int
}

func (mipc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de disco y es una %s", mipc.ram, mipc.disco, mipc.marca)
}

func main() {

	mipc := pc{ram: 16, marca: "msi", disco: 100}

	fmt.Println(mipc)

}
