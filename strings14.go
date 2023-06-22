package main

import (
	"fmt"
	"strconv"
)

//Solicite ao usuário uma string e informe se ela é uma sequência numérica decrescente (exemplo: "987" ou "54321").

func main() {
	str := "98765"
	dec := true
	fmt.Println("String: ", str)
	for i := 0; i < len(str)-1; i++ {
		strconv.Atoi(str)
		if str[i+1] == str[i]-1 {
			dec = true
		} else {
			dec = false
		}
		if !dec {
			break
		}
	}
	if dec {
		fmt.Print("Ordem decrescente")
	} else {
		fmt.Print("Ordem não decrescente")
	}
}
