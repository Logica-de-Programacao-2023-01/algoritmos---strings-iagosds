package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário
func main() {
	var f, x, y string
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&f)
	fmt.Print("Qual letra deseja substituir? ")
	fmt.Scan(&x)
	fmt.Print("Por qual letra deseja substitui-la? ")
	fmt.Scan(&y)
	f = strings.ReplaceAll(f, x, y)
	fmt.Print(f)
}
