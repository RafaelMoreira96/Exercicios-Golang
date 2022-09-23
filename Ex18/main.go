package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func daVolta (n int, min int, max int) int {
	if (n > max) { return  (min + (n - max)) }
}

func cifra(texto string, param rune) string {
	for _, v := range strings.ToLower(texto) {
		if v <= 122 && v >= 97 {
			// Não entendi o que esse if faz, então tirei;
			//if v == 122 {
			//	v = 96
			//}
			
			
			// Como é bom sempre explicar o que tá acontecendo,
			// se você colocar a letra 121 da tabela ASCII e um salto de 13;
			// Você terminará com o caracter å, ASCII 134, ao invés de dar a volta;
			// O código antigo está aqui comentado
			// v = v + (param)
			
			v = daVolta(v + (param), 97, 122)
			fmt.Printf("%c", v)
		}
	}
	return texto
}

func main() {
	fmt.Printf("\n\n -- 18 - A Cifra de César é a técnica mais simples usada pela criptografia. Usada por Júlio César para a transmissão de mensagens secretas, a cifra é baseada na troca de uma letra por outra sempre obedecendo um código (ou melhor, uma chave). Faça um programa com base no que leu a cima que é criptografia utilizando da Cifra de César. --\n\n")
	var param rune
	var texto string
	// carac := []string{" ", "!", "?", ",", ".", ":", ";"}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Digite um texto: ")
	scanner.Scan()
	texto = scanner.Text()

	fmt.Printf("Quantas letras vai pular?: ")
	fmt.Scan(&param)

	for _, v := range strings.ToLower(texto) {
		fmt.Printf("%c", v)
	}

	fmt.Printf("\n\n")

	fmt.Println(cifra(texto, param))
}
