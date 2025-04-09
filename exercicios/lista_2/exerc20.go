package main

import "fmt"

func main() {
	var (
		preconormal, precofinal float64
		condicao string
	)

	fmt.Scan(&preconormal)
	fmt.Scan(&condicao)

	if preconormal < 0 {
		fmt.Println("Numero invalido")
	}

	if condicao == "vista" {
		precofinal = preconormal - (preconormal * 10.0/100.0)

	} else if condicao == "vista_credito" {
		precofinal = preconormal - (preconormal - 10.0/100.0)

	} else if condicao == "em_2_vezes" {
		precofinal = preconormal

	} else if condicao == "em_3_vezes" {
		precofinal = preconormal + (preconormal * 10.0/100.0)

	} else {
		fmt.Println("Condição invalida")
	}

	fmt.Printf("%.2f", precofinal)
}