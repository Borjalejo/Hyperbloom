package main

import "fmt"

func say(text string, c chan string) {
	c <- text

}

func main() {

	c := make(chan string, 1)
	fmt.Println("Aloha")
	go say("Bye Bye", c)

	fmt.Println(<-c)
}
