package main

import "fmt"

//Escreva um programa que receba uma string e verifique se ela é um palíndromo. Informe ao usuário se é ou não.

func main() {
	var p, palindromo string
	fmt.Print("Digite um palavra: ")
	fmt.Scan(&p)
	runes := []rune(p)
	for i := len(runes); i >= 0; i-- {
		palindromo = string(runes[i])
	}
	if palindromo == p {
		fmt.Print("A palavra é um palíndromo")
	} else {
		fmt.Print("Não é um palíndromo")
	}
}
