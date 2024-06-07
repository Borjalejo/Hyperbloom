package main

import "fmt"

func mensaje(text string, c chan string) {
	c <- text
}

func main() {

	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "mensaje 2"

	fmt.Println(len(c), cap(c))

	//Range y Close
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	//Select
	email := make(chan string)
	mail := make(chan string)
	go mensaje("Mensaje 1", email)
	go mensaje("mensaje 1", mail)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("Email recibido de E-mail", m1)
		case m2 := <-mail:
			fmt.Println("Email recibido de Mail", m2)
		}
	}
}
