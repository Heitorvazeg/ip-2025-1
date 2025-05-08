package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n int
	)
	fmt.Scan(&n)
	vetor := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	max := 0.0
	min := 0.0

	for j := 0; j < n; j++ {
		if vetor[j] > max {
			max = vetor[j]
			if min == 0 {
				min = vetor[j]
			}
		}
		if vetor[j] < float64(min) {
			min = vetor[j]
		}
	}
	array := make([]float64, n)

	for l := 0; l < n; l++ {
		norm := (vetor[l] - min) / (max - min)
		if math.IsNaN(norm) {
			norm = 0.0
		}
		array[l] = norm
		fmt.Printf("%.2f ", array[l])
	}
}
