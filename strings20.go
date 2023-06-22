package main

import (
	"fmt"
	"strings"
	"unicode"
)

//Solicite ao usuário uma string e informe se ela é está em camelCase e quantas palavras possuí. CamelCase é quando a
//primeira letra de cada palavra é maiúscula, e as demais são minúsculas. Exemplo: "estaStringEstaEmCamelCase"

func main() {
	str := "String Em Camel Case"
	s := strings.Fields(str)
	camel := true
	for _, i := range s {
		if !unicode.IsUpper(rune(i[0])) {
			camel = false
			break
		}
		for c := 1; c < len(i); c++ {
			if !unicode.IsLower(rune(i[c])) {
				camel = false
				break
			}
		}
	}
	if camel {
		fmt.Println("A string esta em CamelCase")
	} else {
		fmt.Println("A string não está em CamelCase")
	}
}
