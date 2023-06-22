package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Quantas palavras essa string tém?"
	fmt.Println(s)
	p := strings.Fields(s)
	np := len(p)
	fmt.Printf("A string contém %d palavras", np)
}
