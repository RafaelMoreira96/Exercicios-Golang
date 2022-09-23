package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("\n\n -- 17 - Escrever uma lista de nomes que retorne em ordem alfabética -- \n\n")

	nomes := []string{"Tereza", "Rafael", "Brian", "Gustavo", "Mateus", "Aarão", "Alberto"}
	fmt.Println("Original:", nomes)

	sort.Strings(nomes)
	fmt.Println("Organizado:", nomes)
}
