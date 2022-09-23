package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeVogal(texto string) {
	vogais := []string{"a", "à", "á", "ã", "â", "e", "è", "é", "ê", "ẽ", "i", "í", "ì", "ĩ", "î", "o", "ò", "ó", "õ", "ô", "u", "ù", "ú", "ũ", "û"}
	var vogaisUpper []string

	// For para criar um outro vetor de letras maiúsculas a partir de 'vogais'
	for _, v := range vogais {
		v = strings.ToUpper(v)
		vogaisUpper = append(vogaisUpper, v)
	}

	for _, v := range vogais {
		texto = strings.ReplaceAll(texto, v, "")
		for _, v := range vogaisUpper {
			texto = strings.ReplaceAll(texto, v, "")
		}
	}
	fmt.Println("Frase sem as vogais: ", texto)
}

func main() {
	fmt.Printf("\n-- 8 - Crie uma função que retira as vogais de uma frase --\n\n")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Digite uma frase: ")
	scanner.Scan()

	texto := scanner.Text()
	removeVogal(texto)
}
