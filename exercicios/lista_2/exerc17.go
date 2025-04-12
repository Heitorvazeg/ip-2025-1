package main

import "fmt"

func main() {
	var (
		conta     int
		metroscub float32
		tipocons  string
		valor     float32
	)

	fmt.Println("Digite os valores: ")
	fmt.Scanf("%d %f %s", &conta, &metroscub, &tipocons)

	if tipocons == "R" {
		valor = 5 + (0.05 * metroscub)
	}

	if tipocons == "I" {
		valor = 800 + ((metroscub - 100) * 0.04)
		if metroscub <= 100 {
			valor = 800
		}
	}

	if tipocons == "C" {
		valor = 500 + ((metroscub - 80) * 0.25)
		if metroscub <= 80 {
			valor = 500
		}
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f", valor)
}
