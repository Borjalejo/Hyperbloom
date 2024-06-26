package main

import "fmt"

//Ver si una palabra es palindromo
func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}

func main() {
	slice := []string{"Aloha", "que", "haces", "?"}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	isPalindromo("ama")
}
