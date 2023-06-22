package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira
func main() {
	str1 := "alguma coisa escrita no slide"
	str2 := "escrita"
	fmt.Println("String1: ", str1)
	fmt.Println("String2: ", str2)
	if strings.Contains(str1, str2) {
		fmt.Print("É uma substring")
	} else {
		fmt.Print("Não é uma substring")
	}
}
