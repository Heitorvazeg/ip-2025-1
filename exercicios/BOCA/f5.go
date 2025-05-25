package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	matriz := make([][]int, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		for k := 0; k < n; k++ {
			fmt.Scan(&matriz[j][k])
		}
	}

	matrizResposta := make([][]int, n)
	for l := 0; l < n; l++ {
		matrizResposta[l] = make([]int, n)
	}

	for a := 0; a < n; a++ {
		for e := 0; e < n; e++ {
			matrizResposta[a][e] = matriz[n-1-e][a]
		}
	}

	for q := 0; q < n; q++ {
		for w := 0; w < n; w++ {
			fmt.Printf("%d ", matrizResposta[q][w])
		}
		fmt.Println()
	}
}