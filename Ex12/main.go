package main

import "fmt"

func verificaValor(metrica, detractors, passives, promoters int) {
	if metrica <= 6 {
		fmt.Println("\nObrigado pelo seu feedback, vamos melhorar!")
		detractors++
	}

	if metrica == 7 || metrica == 8 {
		fmt.Println("\nObrigado pelo feedback...")
		passives++
	}

	if metrica == 9 || metrica == 10 {
		fmt.Println("\nObrigado pelo feedback positivo!")
		promoters++
	}
}
func metricaNPS(metrica, detractors, passives, promoters, ctrlFor int) {
	var num int
	for num = 0; ctrlFor == 1; num++ {
		fmt.Printf("Qual é o nível do nosso atendimento, sendo '0' como ruim e '10' como excelente: ")
		fmt.Scan(&metrica)

		verificaValor(metrica, detractors, passives, promoters)
		if metrica > 10 {
			fmt.Printf("Valor inválido! Processo encerrado.")
			num--
		}

		fmt.Printf("\nDeseja executar o programa novamente? [1 - Sim | 0 - Não]")
		fmt.Scan(&ctrlFor)
	}
	fmt.Println("\n", num, "pessoa(s) responderam a esta pergunta.")
}

func main() {
	fmt.Printf("\n-- 12 - Crie uma função  que classifique o  nível de satisfação do cliente usando METRICA NPS --\n\n")
	var metrica, detractors, passives, promoters int
	ctrlFor := 1
	metricaNPS(metrica, detractors, passives, promoters, ctrlFor)
}
