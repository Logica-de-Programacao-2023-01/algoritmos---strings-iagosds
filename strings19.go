package main

import (
	"bufio"
	"fmt"
	"os"
)

func invert(any string) string {
	var new string
	var c int = 0
	for i := len(any); i >= 0; i-- {
		new[c] = any[i]
		c++
	}
	return new
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Escreva qualquer frase aleatória: ")
	scanner.Scan()
	frase := scanner.Text()
	invertida := invert(frase)
	fmt.Print("Frase invertida: ", invertida)
}
