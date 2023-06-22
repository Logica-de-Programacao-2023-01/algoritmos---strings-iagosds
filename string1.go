package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Escreva uma palavra: ")
	fmt.Scan(&s)
	r := strings.ToUpper(s)
	fmt.Printf("A string final é: %s", r)
}
