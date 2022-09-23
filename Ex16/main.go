package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\n-- 16 - Escrever algoritmo em Go que receba a data de nascimento do usuário como entrada,  e retorne o dia da semana que o usuário nasceu. --\n\n")
	var dia int
	var mes int
	var ano int

	fmt.Println("Digite a data, com o dia primeiro, depois o mês e depois o ano. Ex.: 25/5/2002")
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	fmt.Scan(&ano)

	agora := time.Date(ano, time.Month(mes), dia, 00, 00, 00, 00, time.Local)
	fmt.Println("O seu nascimento caiu em uma:")

	dias := map[string]string{
		"Sunday": "Domingo",
		"Monday": "Segunda-Feira", 
		"Tuesday": "Terça-Feira",
		"Wednesday": "Quarta-Feira",
		"Thursday" : "Quinta-Feita",
		"Friday": "Sexta-Feira",
		"Saturday": "Sábado"
	}
	
	fmt.Printf("%s\n", map[agora.Weekday().String()])
	
}
