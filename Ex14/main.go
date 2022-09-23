package main

import "fmt"

func calculaIMC(peso, altura float64) float64 {
	return peso / (altura * altura)
}

func classificacaoIMC(imc float64) {
	fmt.Printf("Seu IMC é de %.1f: \nClassificação: ", imc)
	table := map[int]string { 
		18.5: "Abaixo do peso ideal", 
		24.9: "Normal", 
		29.9: "Sobrepeso",
		34.9: "Obesidade Grau I", 
		39.9: "Obesidade Grau II", 
		999999: "Obesidade Grau III"
	}
	
	for key. value := range table {
		if (imc < key) { 
			Printf("%s", value)
			break
		}
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
