package main

import "fmt"

//Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez)

func main() {
	letras := make(map[string]int)
	str := "varias letras juntas"
	for _, i := range str {
		if string(i) != " " {
			letras[string(i)]++
		}
	}
	for c, i := range letras {
		if i == 1 {
			fmt.Print(c)
		}
	}
}
