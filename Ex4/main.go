package main

import "fmt"

func soma(n1, n2 int8) int8 {
	return n1 + n2
}

func main() {
	var n1, n2 int8
	fmt.Printf("\n\n-- Exercício 4: Crie uma função que some duas notas --\n\n")

	fmt.Printf("Usuário, digite nota A e B: ")
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)

	fmt.Println("\n\n A soma é:", soma(n1, n2))
}
