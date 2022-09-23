package main

import "fmt"

func adicao(a, b float64) float64 {
	return a + b
}

func subtracao(a, b float64) float64 {
	return a - b
}

func multiplicacao(a, b float64) float64 {
	return a * b
}

func divisao(a, b float64) float64 {
	return a / b
}

func main() {
	var ctrl = 1
	var i int
	fmt.Printf("\n-- 6 - Uma função que cria uma calculadora simples, digitando o n1, n2 e a operação desejada --\n\n")

	for i = 0; ctrl == 1; i++ {

		var n1, n2, operacao float64
		fmt.Println("Usuário, informe dois números")
		fmt.Scan(&n1)
		fmt.Scan(&n2)

		fmt.Println("Escolha a operação desejada:")
		fmt.Printf("1 - Adição | 2 - Subtração \n3 - Multiplicação | 4 - Divisão\n")

		fmt.Scan(&operacao)

		switch operacao {
		case 1:
			fmt.Println("O resultado de", n1, "+", n2, "é igual à", adicao(n1, n2))

		case 2:
			fmt.Println("O resultado de", n1, "-", n2, "é igual à", subtracao(n1, n2))

		case 3:
			fmt.Println("O resultado de", n1, "*", n2, "é igual à", multiplicacao(n1, n2))

		case 4:
			fmt.Println("O resultado de", n1, "/", n2, "é igual à", divisao(n1, n2))
		}

		fmt.Println("Deseja rodar o programa novamente? [1 - Sim | 0 - Não] ")
		fmt.Scan(&ctrl)

	}
	fmt.Println("Esta calculadora simples fez", i, "operação(ões). Obrigado!")
}
