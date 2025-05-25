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
		vetor[i] = make([]int, m)
	}

	for j := 0; j < n; j++ {
		for k := 0; k < m; k++ {
			fmt.Scan(&valor)
			vetor[j][k] = valor
		}
	}

	vetorR := make([][]int, m)
	for u := 0; u < m; u++ {
		vetorR[u] = make([]int, n)
	}


	for l := 0; l < n; l++ {
		for z := 0; z < m; z++ {
			vetorR[z][l] = vetor[l][z]
		}
	}

	for p := 0; p < m; p++ {
		for a := 0; a < n; a++{
			fmt.Printf("%d ", vetorR[p][a])
		}
		fmt.Print("\n")
	}
}