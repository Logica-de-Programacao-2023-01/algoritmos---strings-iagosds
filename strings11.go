package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.

func main() {
	str := "alguma coisa"
	vogais := []string{"a", "e", "i", "o", "u"}
	fmt.Println("String: ", str)
	for i := range vogais {
		str = strings.ReplaceAll(str, vogais[i], "")
	}
	fmt.Print("Nova string: ", str)
}
