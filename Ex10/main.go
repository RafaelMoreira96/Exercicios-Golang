package main

import "fmt"

func divisivel(n, j int) bool {
	if (n % j) == 0 {
		return true
	}
	return false
}

func primo(n, i int) bool {
	if n == 1 {
		return false
	}
	if n == i {
		return true
	}
	if divisivel(n, i) {
		return false
	}
	return primo(n, i+1)
}

func main() {
	fmt.Printf("\n\n --10 - uma função para dizer se o numero inserido é primo ou n -- \n\n")
	var n int
	//count := 1
	fmt.Printf("Digite um número desejado: ")
	fmt.Scan(&n)

	if primo(n, 2) == true {
		fmt.Println("Número primo")
	} else {
		fmt.Println("Número não é primo")
	}
}
