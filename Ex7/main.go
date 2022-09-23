package main

import "fmt"

func loopDez(media, maxima float64) (float64, float64) {
	var numero float64
	for i := 0; i < 10; i++ {
		fmt.Scan(&numero)
		if numero > maxima {
			maxima = numero
		}
		media += numero
	}
	media = media / 10
	return media, maxima
}

func main() {
	fmt.Printf("\n-- 7 - Crie uma função que receba 10 números, faça a média e mostre o maior número --\n\n")
	var media, maxima float64
	fmt.Println("Digite os 10 numeros abaixo:")
	media, maxima = loopDez(media, maxima)
	fmt.Println("Média:", media)
	fmt.Println("Maior Numero:", maxima)
}
