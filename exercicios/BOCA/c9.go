package main

import (
	"fmt"
	"math"
)

func main() {
	var X, fatorial float64
	fmt.Scan(&X)

	sinal := -1.0
	S := 1.0
	count := 2

	for i := 1; i <= 20; i++ {
		fatorial = 1.0

		for j := 1; j <= count; j++ {
			fatorial *= float64(j)
		}

		termo := sinal * (math.Pow(X, float64(count)) / fatorial)
		count += 2

		S += termo
		sinal *= -1.0
	}

	cos := math.Cos(X)
	dif := S - cos

	fmt.Printf("%.4f %.4f %.4f", S, cos, dif)
}
