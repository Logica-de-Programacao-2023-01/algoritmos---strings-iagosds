package main

import (
	"fmt"
)

//Escreva um programa que receba duas strings e verifique se elas são anagramas. Informe ao usuário se são ou não

func main() {
	cont := make(map[byte]int)
	cont2 := make(map[byte]int)
	var p1, p2 string
	fmt.Print("Digite a primeira palavra: ")
	fmt.Scan(&p1)
	for i := 0; i < len(p1); i++ {
		cont[p1[i]]++
	}
	fmt.Print("Digite a segunda palavra: ")
	fmt.Scan(&p2)
	for i := 0; i < len(p2); i++ {
		cont2[p2[i]]++
	}
	anag := true
	for l, _ := range cont {
		if cont[l] != cont2[l] {
			anag = false
		}
	}
	if anag {
		fmt.Print("São anagramas")
	} else {
		fmt.Print("Não são anagramas")
	}
}
