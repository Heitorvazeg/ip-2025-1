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
	count := 0

	for z := 0; z < n; z++ {
		auxiliar := matriz[z][0]
		for a := 1; a < m; a++ {
			if auxiliar > matriz[z][a] && count < 1{
				fmt.Println("NAO")
				count++
				break
			}
			auxiliar = matriz[z][a]
		}
	}

	if count < 1 {
		for x := 0; x < m; x++ {
			auxiliar := matriz[0][x]
			for e := 1; e < n; e++ {
				if auxiliar > matriz[e][x] && count < 1 {
					fmt.Println("NAO")
					count++
					break
				}
				auxiliar = matriz[e][x]
			}
		}

	}

	if count < 1 {
		fmt.Println("SIM")
	}
}