package main

import (
	"fmt"
)

func main() {
	var horas int
	var valor float32

	fmt.Println("Digite o numero de horas locadas: ")
	fmt.Scan(&horas)

	if horas%3 == 0 {
		valor = float32((horas / 3) * 10)

	} else {
		resto := horas % 3
		valor = float32((horas/3)*10 + (resto * 5))
	}
	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}
