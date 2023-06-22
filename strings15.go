package main

import (
	"fmt"
	"strings"
)

//Solicite ao usuário uma string e substitua todas as vogais por '*' (asterisco)

func main() {
	vog := "aeiouAEIOU"
	runvog := []rune(vog)
	var p string
	fmt.Print("Digite um palavra: ")
	fmt.Scan(&p)
	for i := range runvog {
		p = strings.ReplaceAll(p, string(runvog[i]), "*")
	}
	fmt.Print("Nova string: ", p)
}
