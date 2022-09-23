package main

import "fmt"

func multiplier(n1, n2, n3 int) int {
	return n1 * n2 * n3
}
func main() {
	var n1, n2, n3 int
	fmt.Printf("\n\n-- Exercício 9: Receber três números e faça a multiplicação entre eles --\n\n")

	fmt.Printf("Usuário, insira três números para multiplicação: ")
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	fmt.Scanf("%d", &n3)

	fmt.Println("O resultado da multiplicação é: ", multiplier(n1, n2, n3))
}
