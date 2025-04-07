package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a float64
	var area, volume float64

	fmt.Println("Digite a altura: ")
	fmt.Scan(&h)

	fmt.Println("Digite o valor da aresta: ")
	fmt.Scan(&a)

	raiz := math.Pow(3, 1.0/2.0)
	apow := math.Pow(a, 2)
	area = (3 * apow * raiz) / 2
	volume = (area * h) / 3

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS", volume)
}
