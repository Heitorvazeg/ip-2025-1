package main

import (
	"fmt"
	"math"
)

func main() {
	var entrada int

	for i := 0; ; i++ {
		fmt.Scan(&entrada)

		if entrada <= 0 {
			break

		} else {
			potenciaint := int(math.Pow(float64(entrada), 1.0/2.0))
			potenciafloat := math.Pow(float64(entrada), 1.0/2.0)

			if potenciafloat == float64(potenciaint) {
				fmt.Printf("%d eh quadrado perfeito\n", entrada)

			} else {
				fmt.Printf("%d nao eh quadrado perfeito\n", entrada)
			}
		}
	}
}
