package main

import "fmt"

func converterCtoF(tc float64) {
	f := (tc * 9 / 5) + 32
	fmt.Println(f, "°F")
}

func main() {
	fmt.Printf("\n-- 11 - Crie uma função recebe uma temperatura em graus celsius e converta para Fahrenheit. --\n\n")
	var tc float64

	fmt.Printf("Informe os graus Celsius: ")
	fmt.Scan(&tc)

	converterCtoF(tc)
}
