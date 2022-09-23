package main

import (
	"fmt"

	ex "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
)

func convertToUSD_API(real float64) float64 {
	SwapTest := swap.NewSwap()

	SwapTest.AddExchanger(ex.NewTheMoneyConverterApi(nil)).Build()

	//aqui só pra pegar o valor do preço do dólar em real
	usdRate := SwapTest.Latest("USD/BRL")
	fmt.Println("-- Cotação do dólar de hoje --")
	fmt.Println("US$ 1 = R$ ", usdRate.GetRateValue())

	return real / usdRate.GetRateValue()
}

func convertToUSD_NoAPI(real, manual float64) float64 {
	return real / manual
}

func main() {
	fmt.Printf("\n-- 13 - Criar uma função que converta a moeda Real , para moeda dólar --\n\n")
	var real, total, ctrlSwitchCase, manual float64
	ctrlLoop := 1

	for i := 0; ctrlLoop == 1; i++ {
		fmt.Printf("Digite qual é o valor que você deseja converter: ")
		fmt.Scan(&real)

		fmt.Printf("Deseja utilizar a API de conversão ou deseja informar o valor atual do Dolar? [1 - API | 0 - Inserção Manual]: ")
		fmt.Scan(&ctrlSwitchCase)

		switch ctrlSwitchCase {
		case 1:
			fmt.Printf("Site utilizado: themoneyconverter.com\n\n")
			total = convertToUSD_API(real)
		case 0:
			fmt.Printf("Insira o valor manualmente: [Ex.: 5.23]: ")
			fmt.Scan(&manual)
			total = convertToUSD_NoAPI(real, manual)
		}
		fmt.Println("\n-- Conversão do Real Brasileiro para o Dólar Americano --")
		fmt.Printf("R$ %.2f = US$ %.2f\n\n", real, total)

		fmt.Printf("Deseja fazer outra conversão de outro valor? [1 - Sim | 0 - Não]: ")
		fmt.Scan(&ctrlLoop)
		fmt.Println()
	}
}
