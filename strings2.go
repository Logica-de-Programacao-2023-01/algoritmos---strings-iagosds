package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Essa string contém espaços"
	fmt.Println("Essa é a string: ", s)
	r := strings.ReplaceAll(s, " ", "")
	fmt.Print("A nova string é: ", r)
}
