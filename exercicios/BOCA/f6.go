package main

import "fmt"

func main() {
	var n, m, k int

	fmt.Scan(&n, &k, &m)

	matriz_a := make([][]int, n)
	for i := 0; i < n; i++ {
		matriz_a[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&matriz_a[i][j])
		}
	}

	matriz_b := make([][]int, k)
	for a := 0; a < k; a++ {
		matriz_b[a] = make([]int, m)
		for b := 0; b < m; b++ {
			fmt.Scan(&matriz_b[a][b])
		}
	}

	matrizResposta := make([][]int, n)
	for e := 0; e < n; e++ {
		matrizResposta[e] = make([]int, m)
	}

	for c := 0; c < n; c++ {
		for d := 0; d < m; d++ {
			valor := 0
			for g := 0; g < k; g++ {
				valor += matriz_a[c][g]*matriz_b[g][d]
			}
			matrizResposta[c][d] = valor
		}
	}

	for o := 0; o < n; o++ {
		for p := 0; p < m; p++ {
			fmt.Printf("%d ", matrizResposta[o][p])
		}
		fmt.Println()
	}
}