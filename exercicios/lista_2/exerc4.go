package main

import (
	"fmt"
	"math"
)

func main() {
	var entrada float64
	fmt.Scan(&entrada)

	raiz := math.Pow(entrada, 0.5)

	if raiz >= 0 {
		fmt.Printf("A raiz é: %.2f", raiz)

	} else {
		quadrado := math.Pow(entrada, 2)
		fmt.Printf("O quadrado é: %.2f", quadrado)
	}
}
