package main

import (
	"fmt"
	"math"
)

func main() {
	var X float64
	var fatorial, termo float64
	fmt.Scan(&X)

	for i := 0.0; i <= X+0.00001; i += 0.1 {
		i = math.Round(i*10) / 10

		sinal := -1.0
		count := 3
		seno := i
		fatorial = 1

		for k := 1; k <= 3; k++ {
			fatorial = 1

			for j := 1; j <= count; j++ {
				fatorial *= float64(j)
			}

			termo = sinal * (math.Pow(i, float64(count)) / fatorial)
			count += 2

			seno += termo

			sinal *= -1.0
		}

		fmt.Printf("%.1f ", i)
		fmt.Printf("%.4f\n", seno)
	}
}
