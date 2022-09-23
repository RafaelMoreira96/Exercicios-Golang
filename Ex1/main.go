package main

import "fmt"

func imprimeArvore(matriz [10][7]string) {
	for i := 0; i < 10; i++ {
		fmt.Println()
		for j := 0; j < 7; j++ {
			fmt.Printf("%s ", matriz[i][j])
		}
	}
	fmt.Println()
}

// for i = 0; i < 10; i++ {
// 	fmt.Println(" ", matriz[i])
// }

func main() {
	fmt.Printf("\n-- 1 - criar um laço que crie com asterisco uma arvore de natal --\n\n")

	var matriz [10][7]string
	var d, e int
	d = 4
	e = 2
	i := 1
	count := 0

	matriz[0][3] = "*"
	for i < 15 { // percorrer todas as linhas
		matriz[i][e] = "*"
		matriz[i][d] = "*"
		if d <= 7 && e >= 0 {
			matriz[i][e] = "*"
			matriz[i][d] = "*"
			if e != 0 {
				e--
			}
			d++
		}
		for d == 7 && e == 0 {
			for j := 0; j < 5; j++ {
				matriz[i][j] = "*"
				matriz[i][j+2] = ""
			}
			d = 4
			e = 2
			count++
		}
		i++
		if count == 2 {
			for i <= 9 {
				matriz[i][2] = "*"
				matriz[i][4] = "*"
				i++
			}
			break
		}
	}

	imprimeArvore(matriz)

}
