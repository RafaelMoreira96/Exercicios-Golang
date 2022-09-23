package main

import "fmt"

func convertSectoMin(min, sec int) (int, int) {
	min = min + 0
	sec = sec + 0
	for min = 0; sec >= 60; min++ {
		sec = sec - 60
	}
	return min, sec
}

func main() {
	var min, sec, secTemp int
	fmt.Printf("\n-- Exercício 2: Criar uma função que converta segundos para minutos e segundos. Ex: 80seg = 1 min e 20seg --\n\n")

	fmt.Printf("Usuário, informe segundos: ")
	fmt.Scan(&sec)
	secTemp = sec

	min, sec = convertSectoMin(min, sec)

	fmt.Println(secTemp, "segundos convertidos em minutos e segundos é: ")

	if min >= 2 && sec > 1 {
		fmt.Println(min, "minutos e", sec, "segundos")
	} else if min > 2 && sec == 1 {
		fmt.Printf("%d minutos e %d segundo", min, sec)
	} else if min == 1 && sec > 1 {
		fmt.Println(min, "minuto e", sec, "segundos")
	} else {
		fmt.Println(min, "minuto e", sec, "segundo")
	}
}
