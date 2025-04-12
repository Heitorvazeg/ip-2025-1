package main

import (
	"fmt"
	"math"
)

func main() {
	var X float64
	var fatorial float64
	fmt.Scan(&X)

	sinal := -1.0
	S := X

	for i := 1; i <= 20; i++ {
		fatorial = 1

		for j := 1; j <= i; j++ {
			fatorial *= float64(j)
		}

		termo := math.Pow(X, float64(i)) / fatorial

		S += sinal * termo
		sinal *= -1
	}
	fmt.Printf("%.5f", S)
}
