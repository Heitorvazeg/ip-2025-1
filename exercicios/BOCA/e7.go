package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	vetor_a := make([]int, n)
	vetor_b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor_a[i])
	}
	for j := 0; j < n; j++ {
		fmt.Scan(&vetor_b[j])
	}

	produtoEscalar := 0
	for k := 0; k < n; k++ {
		produtoEscalar += vetor_a[k] * vetor_b[k]
	}
	fmt.Print(produtoEscalar)
}
