package main

import "fmt"

func main() {
	var n, m, valor int

	fmt.Scan(&n, &m)
	matriz := make([][]int, n)

	for i := 0; i < n; i++ {
		matriz[i] = make([]int, m)
	}

	for j := 0; j < n; j++ {
		for k := 0; k < m; k++ {
			fmt.Scan(&valor)
			matriz[j][k] = valor
		}
	}

	for l := 0; l < n; l++ {
		for a := 0; a < m; a++ {
			if a == 0 && (l == 0 || l == n-1)
			|| a == m-1 && (l == 0 || l == n-1) {
				soma := matriz[l+1][a]+matriz[l][a+1]+matriz[l+1][a+1]
				matriz[l][a] = soma

			}
		}
	}

	for z := 0; z < n; z++ {
		for b := 0; b < m; b++ {
			fmt.Printf("%d ", matriz[z][b])
		}
		fmt.Println()
	}
}