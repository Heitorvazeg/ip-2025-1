package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scan(&n)
	vetor := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	s := 0.0
	k := n - 1

	for j := 0; j < n/2; j++ {
		s += math.Pow(vetor[j]-vetor[k], 3)
		k--
	}
	fmt.Printf("%.2f\n", s)
}
