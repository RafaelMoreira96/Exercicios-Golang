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

func plural(n int) string {
	if (n  <= 1) { return "" }
	return "s"
}

func main() {
	var min, sec, secTemp int
	fmt.Printf("\n-- Exercício 2: Criar uma função que converta segundos para minutos e segundos. Ex: 80seg = 1 min e 20seg --\n\n")

	fmt.Printf("Usuário, informe segundos: ")
	fmt.Scan(&sec)
	secTemp = sec
	// A função tá ótima, aqui é só questão de legibilidade mesmo;
	//Trocadilho só por precaução
	if (sec < 0) { fmt.Printf("Não trabalhamos com relatividade, (ainda).") }
	else {
		min, sec = convertSectoMin(min, sec)
		fmt.Printf("%d minuto%s e %d segundo%s", min, plural(min), sec, plural(sec))
	}
}
