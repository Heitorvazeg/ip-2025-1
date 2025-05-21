package main

import "fmt"

func main() {
	var (
		n int
		m int
		valor int
	)

	fmt.Scan(&n, &m)
	vetor := make([][]int, n)

	for i := 0; i < n; i++ {
		vetor_2 := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&valor)
			vetor_2[j] = valor
		}
		vetor[i] = vetor_2
	}

	for k := 0; k < n; k++ {
		somaL := 0
		for l := 0; l < m; l++ {
			somaL += vetor[k][l]
		}
		fmt.Printf("%d ", somaL)
	}
	
	fmt.Println("\n")

	for z := 0; z < m; z++ {
		somaC := 0
		for a := 0; a < n; a++ {
			somaC += vetor[a][z]
		}
		fmt.Printf("%d ", somaC)
	}
}
