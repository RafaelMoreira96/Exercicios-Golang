package main

import "fmt"

func anoEmMinutos(ano int64) int64 {
	var anoEmMinutos int64 = 525600
	return ano * anoEmMinutos
}

func main() {
	fmt.Printf("\n-- 3 - Criar uma função que converta anos para minutos e segundos. --\n\n")

	var ano int64

	fmt.Printf("Usuário, informe a quantidade de anos: ")
	fmt.Scan(&ano)

	resultado := anoEmMinutos(ano)
	fmt.Println("O resultado é: ", resultado)
}
