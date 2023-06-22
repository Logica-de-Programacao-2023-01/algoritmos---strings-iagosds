package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checker(x string) {

	if strings.Contains(x, "9") || strings.Contains(x, "8") || strings.Contains(x, "7") || strings.Contains(x, "6") || strings.Contains(x, "5") || strings.Contains(x, "4") || strings.Contains(x, "3") || strings.Contains(x, "2") || strings.Contains(x, "1") || strings.Contains(x, "0") {
		fmt.Println("A string contém pelo menos um número")
	} else {
		fmt.Println("A string não contém nenhum número")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Escreva qualquer frase aleatória: ")
	scanner.Scan()
	frase := scanner.Text()
	checker(frase)
}
