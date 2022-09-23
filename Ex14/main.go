package main

import "fmt"

func calculaIMC(peso, altura float64) float64 {
	return peso / (altura * altura)
}

func classificacaoIMC(imc float64) {
	fmt.Printf("Seu IMC é de %.1f: \nClassificação: ", imc)
	switch {
	case imc <= 18.5:
		fmt.Println("Abaixo do peso ideal")
	case imc >= 18.6 && imc <= 24.9:
		fmt.Println("Normal")
	case imc >= 25 && imc <= 29.9:
		fmt.Println("Sobrepeso")
	case imc >= 30 && imc <= 34.9:
		fmt.Println("Obesidade Grau I")
	case imc >= 35 && imc <= 39.9:
		fmt.Println("Obesidade Grau II")
	case imc >= 40:
		fmt.Println("Obesidade Grau III")
	}
}
func main() {
	fmt.Printf("\n-- 14 - crie uma função que calcula o IMC --\n\n")
	var imc, peso, altura float64

	fmt.Printf("Digite o peso: ")
	fmt.Scan(&peso)

	fmt.Printf("Digite o altura: ")
	fmt.Scan(&altura)

	imc = calculaIMC(peso, altura)

	classificacaoIMC(imc)
}
