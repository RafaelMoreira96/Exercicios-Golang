package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n-- 15 - Faça um programa que mostra na tela os número de 1 a 100 e print em uma matriz quais deles sao numero de Ouro ? --\n\n")

	fmt.Printf("Números de 1 à 100\n\n")

	var matriz [10][9]int
	a := 1
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			matriz[i][j] = a
			a++
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(matriz[i])
	}

	sFib := []int{1, 1}

	for i := 0; sFib[i] < 50; i++ {
		sFib = append(sFib, (sFib[i] + sFib[i+1]))
	}

	fmt.Printf("\nSequencia de Fibonacci (ou, os números de ouro):\n\n")

	fmt.Println(sFib)

}
