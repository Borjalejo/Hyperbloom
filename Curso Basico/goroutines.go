package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	fmt.Println(text)
	defer wg.Done()

}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Aloha")

	wg.Add(1)
	go say("Mundo", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Bye Bye")

	time.Sleep(time.Second * 1)

}
