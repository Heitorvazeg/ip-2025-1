package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	vetorContrario := make([]int, n)
	count := 0
	for j := n - 1; j >= 0; j-- {
		vetorContrario[count] = vetor[j]
		count++
	}

	for k := 0; k < n; k++ {
		if vetor[k] != vetorContrario[k] {
			fmt.Println("NAO")
			return
		}
	}
	fmt.Println("SIM")
}
