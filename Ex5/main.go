package main

import "fmt"

func mediaNota(media float32) float32 {
	return media / 2
}

func main() {
	fmt.Printf("\n-- 5 - Criar uma função que faça média entre notas --\n\n")
	var nota, media float32
	fmt.Printf("Digite duas notas: ")
	for i := 0; i < 2; i++ {
		fmt.Scan(&nota)
		media += nota
	}
	fmt.Printf("A média é %.2f\n", mediaNota(media))
}
