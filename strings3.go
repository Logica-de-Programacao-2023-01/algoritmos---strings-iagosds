package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Essa string sera alterada em por você"
	var carac, subs string
	fmt.Println("String: ", s)
	fmt.Print("Qual caractere deseja substituir? ")
	fmt.Scanln(&carac)
	fmt.Print("Por qual caractere deseja substituílo? ")
	fmt.Scanln(&subs)
	r := strings.ReplaceAll(s, carac, subs)
	fmt.Print("A nova string é: ", r)
}
