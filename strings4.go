package main

import (
	"fmt"
)

// Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.
func main() {
	var p1, p2 string
	fmt.Print("Escreva uma palavra: ")
	fmt.Scan(&p1)
	fmt.Print("Digite outra palavra, igual ou diferente: ")
	fmt.Scan(&p2)
	if p1 == p2 {
		fmt.Print("Palavras são iguais!")
	} else {
		fmt.Print("Palavras são diferentes!")
	}
}
