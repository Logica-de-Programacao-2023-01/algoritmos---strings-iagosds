package main

import "fmt"

func main() {
	var vetor []int
	fmt.Println("Escreva os números")
	for c := 0; c < 4; c++ {
		fmt.Scan(&vetor[c])
	}
	ordem(vetor)
}
func ordem(x []int) {
	dir := 0
	ord := 0

	if x[0] > x[1] {
		dir = 1
	} else {
		dir = 2
	}

	for i := 2; i < len(x); i++ {
		if x[i+1] < x[i] {
			if dir == 2 {
				ord = 3
			}
			ord = 1
		}
		if x[i+1] > x[i] {
			if dir == 1 {
				ord = 3
			}
			ord = 2
		}
	}

	if ord == 1 {
		fmt.Println("Decrescente")
	} else if ord == 2 {
		fmt.Println("Crescente")
	} else {
		fmt.Println("ordem aleatória")
	}
}
