package main

import "fmt"

// Escreva um programa que receba uma string e inverta a ordem dos caracteres. Informe ao usuário o resultado
func main() {
	x := "inversao"
	var nova []rune
	fmt.Println(x)
	for i := len(x); i >= 0; i-- {
		for n := 0; n <= len(x); n++ {
			nova[n] = rune(x[i])
			break
		}
	}
	fmt.Println(nova)
}
