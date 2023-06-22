package main

import (
	"fmt"
	"strconv"
)

//Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante.
//Informe ao usuário se é ou não.

func main() {
	var x string
	fmt.Println("Escreva uma string: ")
	fmt.Scan(&x)
	_, err := strconv.ParseFloat(x, 1)
	if err == nil {
		fmt.Println("A string pode ser um ponto flutuante")
	} else {
		fmt.Println("A string não pode ser um ponto flutuante")
	}
}
